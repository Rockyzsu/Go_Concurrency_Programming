using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace ChatClient
{
    public partial class FrmChat : Form
    {
        public FrmChat()
        {
            InitializeComponent();
        }
        System.Windows.Forms.Timer CheckReceiveTimer = new System.Windows.Forms.Timer();
        TCPClient tcpClient;
        private void FrmChat_Load(object sender, EventArgs e)
        {
            tcpClient = new TCPClient(7777);
            CheckReceiveTimer.Interval = 1000;
        
            CheckReceiveTimer.Tick += CheckReceiveTimer_Tick;
        }

        private void CheckReceiveTimer_Tick(object sender, EventArgs e)
        {
            try
            {
                ThreadStart thStart = new ThreadStart(Pro);//threadStart委托 
                Thread thread = new Thread(thStart);
                thread.Priority = ThreadPriority.Highest;
                thread.IsBackground = true; //关闭窗体继续执行
                thread.Start();
            }
            catch(Exception ex)
            {
                Console.WriteLine(ex.Message);
            }
           
        }
        public void Pro()
        {
            string msg = tcpClient.GetMsg();
            if (this.txtLog.InvokeRequired)
            {
               
                txtLog.Invoke(new Action<TextBox, string>(SetTxtValue), txtLog, msg);
            }
            else
            {
                txtLog.Text +=msg + "\r\n";
            }
          
        }

        private void SetTxtValue(TextBox txt, string value)
        {
            txt.Text += value+"\r\n";
        }

        private void btnSend_Click(object sender, EventArgs e)
        {
            try
            {
                tcpClient.SendMsg(this.txtMsg.Text);
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
        }
        //用户名登录
        private void btnDial_Click(object sender, EventArgs e)
        {
            try
            {
                tcpClient.Login(this.txtUserName.Text);
                this.txtLog.Text ="["+this.txtUserName.Text+ "]连接成功\r\n";
                CheckReceiveTimer.Enabled = true;
                this.btnDial.Enabled = false;
                this.txtUserName.ReadOnly = true;
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
        }

        private void FrmChat_FormClosing(object sender, FormClosingEventArgs e)
        {
            tcpClient.Close();
        }
    }
}
