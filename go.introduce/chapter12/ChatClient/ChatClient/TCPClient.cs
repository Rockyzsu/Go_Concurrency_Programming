using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;

namespace ChatClient
{
    class TCPClient
    {
        private Socket _ClientSocket;                       //客户端通讯套接字
        private IPEndPoint SeverEndPoint;                   //连接到服务器端IP和端口

        public TCPClient(int PORT=7777)
        {
            //服务器通信地址
            SeverEndPoint = new IPEndPoint(IPAddress.Parse("127.0.0.1"), PORT);
            //建立客户端Socket
            _ClientSocket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);

            try
            {
                _ClientSocket.Connect(SeverEndPoint);
                Console.WriteLine("连接服务器端成功！");
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
            }
          
        }

        
        public void SendMsg(string strMsg)
        {
            //注意：服务会将消息自动移除2个Byte
            Byte[] byeArray = Encoding.UTF8.GetBytes(strMsg+"\r\n");
                //发送数据
                _ClientSocket.Send(byeArray);    
        }
        public void Login(string userName)
        {
            //注意：服务会将消息自动移除2个Byte
            Byte[] byeArray = Encoding.UTF8.GetBytes(userName);
            //发送数据
            _ClientSocket.Send(byeArray);
        }
        public string GetMsg()
        {
            byte[] data = new byte[1024];
            //传递一个byte数组，用于接收数据。length表示接收了多少字节的数据
            int length = _ClientSocket.Receive(data);
            string message = Encoding.UTF8.GetString(data, 0, length);//只将接收到的数据进行转化
            return message;
        }
        public void Close()
        {
            try
            {
                //关闭连接
                _ClientSocket.Shutdown(SocketShutdown.Both);
                //清理连接资源
                _ClientSocket.Close();
            }
            catch(Exception ex)
            {
                Console.WriteLine(ex.Message);
            }
        }
    }
}
