package service

import (
	"fmt"
	"github.com/kjk/betterguid"
	"gorm.io/gorm"
	"log"
	"my-streaming-service/dao"
	"net"
	"os"
	"sync"
	// 引入joy4
	"github.com/nareix/joy4/av"
	"github.com/nareix/joy4/format/flv"
	"github.com/nareix/joy4/format/rtmp"
)

// UserStreamServer 用于管理每个用户的RTMP服务器
type UserStreamServer struct {
	Port       int
	Stream     chan av.Packet
	Server     *rtmp.Server
	RecordFile *os.File // 新增字段，用于录制文件
}

// Manager 管理所有用户的RTMP服务器
type Manager struct {
	Servers map[string]*UserStreamServer
	Mutex   sync.Mutex
	Port    int
}

// NewManager 创建一个新的Manager init时使用
func NewManager() *Manager {
	return &Manager{
		Servers: make(map[string]*UserStreamServer),
		Port:    1935, // 起始端口，可以根据需要更改
	}
}

// StartUserStreamServer 启动一个新的RTMP服务器将用户的推流推流至直播间，并返回URL
func (m *Manager) StartUserStreamServer(userID string, roomID string) string {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	// 为新的RTMP服务器选择一个端口
	port := m.Port
	m.Port++

	// 创建RTMP服务器
	server := &rtmp.Server{}

	// 创建录播文件
	recordFilePath := fmt.Sprintf("%s_%s.flv", userID, roomID)
	recordFile, err := os.Create(recordFilePath)
	if err != nil {
		log.Fatalf("无法创建录播文件: %v", err)
	}

	// 初始化UserStreamServer
	userStreamServer := &UserStreamServer{
		Port:       port,
		Stream:     make(chan av.Packet),
		Server:     server,
		RecordFile: recordFile,
	}

	// 绑定RTMP服务器的HandlePublish函数
	server.HandlePublish = func(conn *rtmp.Conn) {
		streams, _ := conn.Streams()

		// 创建FLV Muxer
		muxer := flv.NewMuxer(recordFile)
		err = muxer.WriteHeader(streams)
		if err != nil {
			log.Fatalf("无法写入FLV头: %v", err)
		}

		for {
			packet, err := conn.ReadPacket()
			if err != nil {
				break
			}
			userStreamServer.Stream <- packet

			// 将数据包写入FLV文件进行录播
			err = muxer.WritePacket(packet)
			if err != nil {
				log.Printf("无法写入数据包: %v", err)
				break
			}
		}

		// 关闭FLV Muxer和录播文件

		if err != nil {
			log.Printf("无法关闭FLV Muxer: %v", err)
		}
		err := userStreamServer.RecordFile.Close()
		if err != nil {
			log.Fatalf("无法关闭录播文件: %v", err)
			return
		}
	}

	// 启动RTMP服务器
	go func() {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("无法监听端口 %d: %v", port, err)
		}
		defer func(listener net.Listener) {
			err := listener.Close()
			if err != nil {

			}
		}(listener)
		log.Fatal(server.ListenAndServe())
	}()

	// 将新的UserStreamServer添加到Manager的Servers中
	m.Servers[roomID] = userStreamServer

	// 返回RTMP推流的URL
	return fmt.Sprintf("rtmp://localhost:%d/live/%s", port, roomID)
}

// PushStreamService 推流至直播间
func (m *Manager) PushStreamService(roomID string, packet av.Packet) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	userStreamServer, ok := m.Servers[roomID]
	if !ok {
		log.Printf("没有找到直播间ID为 %s 的RTMP服务器", roomID)
		return
	}

	userStreamServer.Stream <- packet
}

// CreateLiveRoomService 创建直播间
func CreateLiveRoomService(roomName string, userID string) {
	roomId := betterguid.New()
	err := dao.CreateLiveRoom(roomName, userID, roomId)
	if err != nil {
		log.Printf("创建直播间失败: %v", err)
		return
	}
}

// QueryLiveRoomService 查询直播间
func QueryLiveRoomService(roomName string, roomId string) *gorm.DB {
	room, err := dao.QueryLiveRoom(roomName, roomId)
	if err != nil {
		log.Printf("查询直播间失败: %v", err)
		return nil
	}
	return room
}

// DeleteLiveRoomService 删除直播间
func DeleteLiveRoomService(roomId string) bool {
	return dao.DeleteLiveRoom(roomId)
}

// StopUserStreamServer 停止RTMP服务器
func (m *Manager) StopUserStreamServer(roomID string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	userStreamServer, ok := m.Servers[roomID]
	if ok {
		close(userStreamServer.Stream)
		err := userStreamServer.RecordFile.Close()
		if err != nil {
			log.Fatalf("无法关闭录播文件: %v", err)
			return
		}
		delete(m.Servers, roomID)
	} else {
		log.Printf("没有找到直播间ID为 %s 的RTMP服务器", roomID)
	}
}
