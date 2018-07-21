/* 
 * WRANING: These codes below is far away from bugs with the god and his animal protecting
 *                  _oo0oo_                   ┏┓　　　┏┓
 *                 o8888888o                ┏┛┻━━━┛┻┓
 *                 88" . "88                ┃　　　　　　　┃ 　
 *                 (| -_- |)                ┃　　　━　　　┃
 *                 0\  =  /0                ┃　┳┛　┗┳　┃
 *               ___/`---'\___              ┃　　　　　　　┃
 *             .' \\|     |# '.             ┃　　　┻　　　┃
 *            / \\|||  :  |||# \            ┃　　　　　　　┃
 *           / _||||| -:- |||||- \          ┗━┓　　　┏━┛
 *          |   | \\\  -  #/ |   |          　　┃　　　┃神兽保佑
 *          | \_|  ''\---/''  |_/ |         　　┃　　　┃永无BUG
 *          \  .-\__  '-'  ___/-. /         　　┃　　　┗━━━┓
 *        ___'. .'  /--.--\  `. .'___       　　┃　　　　　　　┣┓
 *     ."" '<  `.___\_<|>_/___.' >' "".     　　┃　　　　　　　┏┛
 *    | | :  `- \`.;`\ _ /`;.`/ - ` : | |   　　┗┓┓┏━┳┓┏┛
 *    \  \ `_.   \_ __\ /__ _/   .-` /  /   　　　┃┫┫　┃┫┫
 *=====`-.____`.___ \_____/___.-`___.-'=====　　　┗┻┛　┗┻┛ 
 *                  `=---='　　　
 *          佛祖保佑       永无BUG
 */
// =============================================================================== 
// Author              :    Frankie.W
// Create Time         :    2018/3/14 11:50:07
// Update Time         :    2018/3/14 11:50:07
// Class Version       :    v1.0.0.0
// Class Description   :    
// ===============================================================================
using MySql.Data.MySqlClient;
using System;
using System.Collections.Generic;
using System.IO;
using System.Text.RegularExpressions;
using System.Web.Script.Serialization;
using System.Windows.Forms;
// ===============================================================================
namespace TrojanCommandSender.Src
{
    class MySQLConnector
    {
        private MySqlConnection m_Connection;
        private ServerConfig m_Config;

        public MySQLConnector(string strConfigFile)
        {
            m_Connection = null;
            m_Config = new ServerConfig();

            string configFilePath = Directory.GetCurrentDirectory() + "\\" + strConfigFile;
            if (!Init(configFilePath))
            {
                MessageBox.Show("加载配置文件: " + configFilePath + " 失败!请检查文件是否存在以及格式是否正确.");
            }
        }

        private bool Init(string strConfigFile)
        {
            try
            {
                if (!File.Exists(strConfigFile))
                {
                    Program.mainForm.AppendTextToTextbox("配置文件: " + strConfigFile + " 不存在!", CONFIG.ERROR_COLOR);
                    return false;
                }
                using (StreamReader r = new StreamReader(strConfigFile))
                {
                    string jsonFile = r.ReadToEnd();
                    var jss = new JavaScriptSerializer();
                    m_Config = jss.Deserialize<ServerConfig>(jsonFile);
                    m_Config.mysql_host = GetIPFromStrings(m_Config.mysql_host);
                    if (!m_Config.IsValid())
                    {
                        Program.mainForm.AppendTextToTextbox("配置文件: " + strConfigFile + " 加载出错!", CONFIG.ERROR_COLOR);
                        return false;
                    }
                    else
                    {
                        Program.mainForm.AppendTextToTextbox("即将访问数据库: " + m_Config.mysql_host + 
                            " 表名:" + m_Config.mysql_name + " 账户:" + m_Config.mysql_user, CONFIG.INFO_COLOR);
                    }
                }

                string connStr = "server=" + m_Config.mysql_host + ";uid=" + m_Config.mysql_user +
                    ";pwd=" + m_Config.mysql_pass + ";database=" + m_Config.mysql_name;
                m_Connection = new MySqlConnection(connStr);
                if(m_Connection == null)
                {
                    Program.mainForm.AppendTextToTextbox("连接数据库失败，连接器未成功创建...", CONFIG.ERROR_COLOR);
                    return false;
                }
            }
            catch(Exception e)
            {
                return false;
            }
            return true;
        }
        private bool OpenConnection()
        {
            try
            {
                if (m_Connection == null)
                {
                    return false;
                }
                m_Connection.Open();
            }
            catch (MySqlException ex)
            {
                switch (ex.Number)
                {
                    case 0:
                        MessageBox.Show("无法连接到数据库...");
                        break;

                    case 1045:
                        MessageBox.Show("账号/密码错误，请重试...");
                        break;
                }
                Program.mainForm.AppendTextToTextbox("连接数据库失败：" + ex.ToString(), CONFIG.ERROR_COLOR);
                return false;
            }
            return true;
        }
        private bool CloseConnection()
        {
            try
            {
                if (m_Connection == null)
                {
                    return false;
                }
                m_Connection.Close();
            }
            catch(MySqlException ex)
            {
                Program.mainForm.AppendTextToTextbox("关闭数据库失败：" + ex.ToString(), CONFIG.ERROR_COLOR);
                return false;
            }
            return true;
        }
        private string GetIPFromStrings(string str)
        {
            Match m = Regex.Match(str, @"\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}");
            if (m.Success)
            {
                return m.Value;
            }
            return string.Empty;
        }
        public bool TestConnection()
        {
            try
            {
                // 直接尝试一次连接
                if (!OpenConnection())
                {
                    Program.mainForm.AppendTextToTextbox("测试连接数据库失败...", CONFIG.ERROR_COLOR);
                    return false;
                }
                // 马上关闭
                if (!CloseConnection())
                {
                    Program.mainForm.AppendTextToTextbox("测试连接数据库失败...", CONFIG.ERROR_COLOR);
                    return false;
                }
            }
            catch (Exception ex)
            {
                Program.mainForm.AppendTextToTextbox("测试连接数据库失败..." + ex.ToString(), CONFIG.ERROR_COLOR);
                return false;
            }
            return true;
        }
        public bool Insert(string query)
        {
            // 增加try catch 防止整个系统抛异常崩溃
            try
            {

                if (!OpenConnection())
                {
                    return false;
                }
                //query.Replace("\"", "`");
                MySqlCommand cmd = new MySqlCommand(query, m_Connection);
               
                if (cmd.ExecuteNonQuery() < 0)
                {
                    return false;
                }
                if (!CloseConnection())
                {
                    return false;
                }
                return true;
            }
            catch {
                return false;
            }
        }
        public List<ClientStruct> GetClientInfos()
        {
            List<ClientStruct> result = new List<ClientStruct>();
            try
            {
                string query = "SELECT * FROM " + CONFIG.CLIENTS_TABLE_NAME;
                if (!OpenConnection())
                {
                    return result;
                }
                MySqlCommand cmd = new MySqlCommand(query, m_Connection);
                MySqlDataReader dataReader = cmd.ExecuteReader();
                while (dataReader.Read())
                {
                    ClientStruct s = new ClientStruct();
                    s.clientGuid = dataReader["guid"] + "";
                    s.clientIP = dataReader["ip"] + "";
                    result.Add(s);
                }
                dataReader.Close();
                if (!CloseConnection())
                {
                    return result;
                }
            }
            catch{}
            return result;
        }
    }
}