namespace TrojanCommandSender
{
    partial class MainForm
    {
        /// <summary>
        /// 必需的设计器变量。
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// 清理所有正在使用的资源。
        /// </summary>
        /// <param name="disposing">如果应释放托管资源，为 true；否则为 false。</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows 窗体设计器生成的代码

        /// <summary>
        /// 设计器支持所需的方法 - 不要修改
        /// 使用代码编辑器修改此方法的内容。
        /// </summary>
        private void InitializeComponent()
        {
            this.components = new System.ComponentModel.Container();
            System.ComponentModel.ComponentResourceManager resources = new System.ComponentModel.ComponentResourceManager(typeof(MainForm));
            this.splitContainer2 = new System.Windows.Forms.SplitContainer();
            this.MainTabControl = new System.Windows.Forms.TabControl();
            this.CommandSendTabPage = new System.Windows.Forms.TabPage();
            this.panel1 = new System.Windows.Forms.Panel();
            this.groupBox2 = new System.Windows.Forms.GroupBox();
            this.panel2 = new System.Windows.Forms.Panel();
            this.SendCommandBtn = new System.Windows.Forms.Button();
            this.AttrPanel = new System.Windows.Forms.Panel();
            this.CommandListView = new System.Windows.Forms.ListView();
            this.UpDownloadTabPage = new System.Windows.Forms.TabPage();
            this.splitContainer4 = new System.Windows.Forms.SplitContainer();
            this.LocalGroupBox = new System.Windows.Forms.GroupBox();
            this.UploadToClientBtn = new System.Windows.Forms.Button();
            this.panel6 = new System.Windows.Forms.Panel();
            this.ChooseLocalFileButton = new System.Windows.Forms.Button();
            this.LocalFileTextBox = new System.Windows.Forms.TextBox();
            this.panel5 = new System.Windows.Forms.Panel();
            this.splitContainer5 = new System.Windows.Forms.SplitContainer();
            this.directoryTree = new System.Windows.Forms.TreeView();
            this.directoryIcons = new System.Windows.Forms.ImageList(this.components);
            this.filesList = new System.Windows.Forms.ListView();
            this.filesIcons = new System.Windows.Forms.ImageList(this.components);
            this.RomateGroupBox = new System.Windows.Forms.GroupBox();
            this.DownloadToLocalBtn = new System.Windows.Forms.Button();
            this.panel8 = new System.Windows.Forms.Panel();
            this.RomateFileTextBox = new System.Windows.Forms.TextBox();
            this.panel7 = new System.Windows.Forms.Panel();
            this.panel4 = new System.Windows.Forms.Panel();
            this.RomateUpdownGuidCB = new System.Windows.Forms.ComboBox();
            this.ClientDetailtabPage = new System.Windows.Forms.TabPage();
            this.OutputRichTextBox = new System.Windows.Forms.RichTextBox();
            this.splitContainer3 = new System.Windows.Forms.SplitContainer();
            this.panel3 = new System.Windows.Forms.Panel();
            this.groupBox1 = new System.Windows.Forms.GroupBox();
            this.ClientListView = new System.Windows.Forms.ListView();
            this.ClientStatus = new System.Windows.Forms.Panel();
            this.splitContainer1 = new System.Windows.Forms.SplitContainer();
            this.UpdateClientsInfoTimer = new System.Windows.Forms.Timer(this.components);
            this.UpdateCommandTimer = new System.Windows.Forms.Timer(this.components);
            ((System.ComponentModel.ISupportInitialize)(this.splitContainer2)).BeginInit();
            this.splitContainer2.Panel1.SuspendLayout();
            this.splitContainer2.Panel2.SuspendLayout();
            this.splitContainer2.SuspendLayout();
            this.MainTabControl.SuspendLayout();
            this.CommandSendTabPage.SuspendLayout();
            this.panel1.SuspendLayout();
            this.groupBox2.SuspendLayout();
            this.panel2.SuspendLayout();
            this.UpDownloadTabPage.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.splitContainer4)).BeginInit();
            this.splitContainer4.Panel1.SuspendLayout();
            this.splitContainer4.Panel2.SuspendLayout();
            this.splitContainer4.SuspendLayout();
            this.LocalGroupBox.SuspendLayout();
            this.panel6.SuspendLayout();
            this.panel5.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.splitContainer5)).BeginInit();
            this.splitContainer5.Panel1.SuspendLayout();
            this.splitContainer5.Panel2.SuspendLayout();
            this.splitContainer5.SuspendLayout();
            this.RomateGroupBox.SuspendLayout();
            this.panel8.SuspendLayout();
            this.panel4.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.splitContainer3)).BeginInit();
            this.splitContainer3.Panel1.SuspendLayout();
            this.splitContainer3.Panel2.SuspendLayout();
            this.splitContainer3.SuspendLayout();
            this.panel3.SuspendLayout();
            this.groupBox1.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.splitContainer1)).BeginInit();
            this.splitContainer1.Panel1.SuspendLayout();
            this.splitContainer1.Panel2.SuspendLayout();
            this.splitContainer1.SuspendLayout();
            this.SuspendLayout();
            // 
            // splitContainer2
            // 
            this.splitContainer2.BorderStyle = System.Windows.Forms.BorderStyle.Fixed3D;
            this.splitContainer2.Dock = System.Windows.Forms.DockStyle.Fill;
            this.splitContainer2.Location = new System.Drawing.Point(0, 0);
            this.splitContainer2.Name = "splitContainer2";
            this.splitContainer2.Orientation = System.Windows.Forms.Orientation.Horizontal;
            // 
            // splitContainer2.Panel1
            // 
            this.splitContainer2.Panel1.Controls.Add(this.MainTabControl);
            this.splitContainer2.Panel1.RightToLeft = System.Windows.Forms.RightToLeft.No;
            // 
            // splitContainer2.Panel2
            // 
            this.splitContainer2.Panel2.Controls.Add(this.OutputRichTextBox);
            this.splitContainer2.Panel2.RightToLeft = System.Windows.Forms.RightToLeft.No;
            this.splitContainer2.RightToLeft = System.Windows.Forms.RightToLeft.No;
            this.splitContainer2.Size = new System.Drawing.Size(937, 730);
            this.splitContainer2.SplitterDistance = 534;
            this.splitContainer2.TabIndex = 1;
            // 
            // MainTabControl
            // 
            this.MainTabControl.Appearance = System.Windows.Forms.TabAppearance.FlatButtons;
            this.MainTabControl.Controls.Add(this.CommandSendTabPage);
            this.MainTabControl.Controls.Add(this.UpDownloadTabPage);
            this.MainTabControl.Controls.Add(this.ClientDetailtabPage);
            this.MainTabControl.Dock = System.Windows.Forms.DockStyle.Fill;
            this.MainTabControl.HotTrack = true;
            this.MainTabControl.Location = new System.Drawing.Point(0, 0);
            this.MainTabControl.Name = "MainTabControl";
            this.MainTabControl.SelectedIndex = 0;
            this.MainTabControl.Size = new System.Drawing.Size(933, 530);
            this.MainTabControl.TabIndex = 0;
            // 
            // CommandSendTabPage
            // 
            this.CommandSendTabPage.Controls.Add(this.panel1);
            this.CommandSendTabPage.Location = new System.Drawing.Point(4, 25);
            this.CommandSendTabPage.Name = "CommandSendTabPage";
            this.CommandSendTabPage.Padding = new System.Windows.Forms.Padding(3);
            this.CommandSendTabPage.Size = new System.Drawing.Size(925, 501);
            this.CommandSendTabPage.TabIndex = 0;
            this.CommandSendTabPage.Text = "命令发送";
            this.CommandSendTabPage.UseVisualStyleBackColor = true;
            // 
            // panel1
            // 
            this.panel1.Controls.Add(this.groupBox2);
            this.panel1.Dock = System.Windows.Forms.DockStyle.Fill;
            this.panel1.Location = new System.Drawing.Point(3, 3);
            this.panel1.Name = "panel1";
            this.panel1.Size = new System.Drawing.Size(919, 495);
            this.panel1.TabIndex = 2;
            // 
            // groupBox2
            // 
            this.groupBox2.Controls.Add(this.panel2);
            this.groupBox2.Controls.Add(this.CommandListView);
            this.groupBox2.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox2.Location = new System.Drawing.Point(0, 0);
            this.groupBox2.Name = "groupBox2";
            this.groupBox2.Size = new System.Drawing.Size(919, 495);
            this.groupBox2.TabIndex = 0;
            this.groupBox2.TabStop = false;
            this.groupBox2.Text = "命令列表";
            // 
            // panel2
            // 
            this.panel2.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.panel2.BackColor = System.Drawing.SystemColors.Control;
            this.panel2.Controls.Add(this.SendCommandBtn);
            this.panel2.Controls.Add(this.AttrPanel);
            this.panel2.Location = new System.Drawing.Point(274, 20);
            this.panel2.Name = "panel2";
            this.panel2.Size = new System.Drawing.Size(639, 469);
            this.panel2.TabIndex = 3;
            // 
            // SendCommandBtn
            // 
            this.SendCommandBtn.Dock = System.Windows.Forms.DockStyle.Bottom;
            this.SendCommandBtn.Location = new System.Drawing.Point(0, 419);
            this.SendCommandBtn.Name = "SendCommandBtn";
            this.SendCommandBtn.Size = new System.Drawing.Size(639, 50);
            this.SendCommandBtn.TabIndex = 0;
            this.SendCommandBtn.Text = "发送命令";
            this.SendCommandBtn.UseVisualStyleBackColor = true;
            this.SendCommandBtn.Click += new System.EventHandler(this.SendCommandBtn_Click);
            // 
            // AttrPanel
            // 
            this.AttrPanel.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.AttrPanel.AutoScroll = true;
            this.AttrPanel.BackColor = System.Drawing.SystemColors.Control;
            this.AttrPanel.BorderStyle = System.Windows.Forms.BorderStyle.Fixed3D;
            this.AttrPanel.Location = new System.Drawing.Point(6, 3);
            this.AttrPanel.Name = "AttrPanel";
            this.AttrPanel.Size = new System.Drawing.Size(627, 410);
            this.AttrPanel.TabIndex = 1;
            // 
            // CommandListView
            // 
            this.CommandListView.Activation = System.Windows.Forms.ItemActivation.OneClick;
            this.CommandListView.Anchor = ((System.Windows.Forms.AnchorStyles)(((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left)));
            this.CommandListView.Location = new System.Drawing.Point(6, 20);
            this.CommandListView.Name = "CommandListView";
            this.CommandListView.Size = new System.Drawing.Size(265, 469);
            this.CommandListView.TabIndex = 0;
            this.CommandListView.UseCompatibleStateImageBehavior = false;
            this.CommandListView.Click += new System.EventHandler(this.CommandListView_Click);
            // 
            // UpDownloadTabPage
            // 
            this.UpDownloadTabPage.Controls.Add(this.splitContainer4);
            this.UpDownloadTabPage.Location = new System.Drawing.Point(4, 25);
            this.UpDownloadTabPage.Name = "UpDownloadTabPage";
            this.UpDownloadTabPage.Padding = new System.Windows.Forms.Padding(3);
            this.UpDownloadTabPage.Size = new System.Drawing.Size(925, 501);
            this.UpDownloadTabPage.TabIndex = 1;
            this.UpDownloadTabPage.Text = "上传下载";
            this.UpDownloadTabPage.UseVisualStyleBackColor = true;
            // 
            // splitContainer4
            // 
            this.splitContainer4.Dock = System.Windows.Forms.DockStyle.Fill;
            this.splitContainer4.Location = new System.Drawing.Point(3, 3);
            this.splitContainer4.Name = "splitContainer4";
            // 
            // splitContainer4.Panel1
            // 
            this.splitContainer4.Panel1.Controls.Add(this.LocalGroupBox);
            // 
            // splitContainer4.Panel2
            // 
            this.splitContainer4.Panel2.Controls.Add(this.RomateGroupBox);
            this.splitContainer4.Size = new System.Drawing.Size(919, 495);
            this.splitContainer4.SplitterDistance = 458;
            this.splitContainer4.TabIndex = 0;
            // 
            // LocalGroupBox
            // 
            this.LocalGroupBox.BackColor = System.Drawing.SystemColors.Control;
            this.LocalGroupBox.Controls.Add(this.UploadToClientBtn);
            this.LocalGroupBox.Controls.Add(this.panel6);
            this.LocalGroupBox.Controls.Add(this.panel5);
            this.LocalGroupBox.Dock = System.Windows.Forms.DockStyle.Fill;
            this.LocalGroupBox.Location = new System.Drawing.Point(0, 0);
            this.LocalGroupBox.Name = "LocalGroupBox";
            this.LocalGroupBox.Size = new System.Drawing.Size(458, 495);
            this.LocalGroupBox.TabIndex = 0;
            this.LocalGroupBox.TabStop = false;
            this.LocalGroupBox.Text = "本地";
            // 
            // UploadToClientBtn
            // 
            this.UploadToClientBtn.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.UploadToClientBtn.Font = new System.Drawing.Font("微软雅黑", 14.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.UploadToClientBtn.Location = new System.Drawing.Point(127, 453);
            this.UploadToClientBtn.Name = "UploadToClientBtn";
            this.UploadToClientBtn.Size = new System.Drawing.Size(185, 39);
            this.UploadToClientBtn.TabIndex = 4;
            this.UploadToClientBtn.Text = "本地->受控端";
            this.UploadToClientBtn.UseVisualStyleBackColor = true;
            this.UploadToClientBtn.Click += new System.EventHandler(this.UploadToClientBtn_Click);
            // 
            // panel6
            // 
            this.panel6.BorderStyle = System.Windows.Forms.BorderStyle.FixedSingle;
            this.panel6.Controls.Add(this.ChooseLocalFileButton);
            this.panel6.Controls.Add(this.LocalFileTextBox);
            this.panel6.Dock = System.Windows.Forms.DockStyle.Top;
            this.panel6.Location = new System.Drawing.Point(3, 411);
            this.panel6.Name = "panel6";
            this.panel6.Size = new System.Drawing.Size(452, 33);
            this.panel6.TabIndex = 3;
            // 
            // ChooseLocalFileButton
            // 
            this.ChooseLocalFileButton.Dock = System.Windows.Forms.DockStyle.Right;
            this.ChooseLocalFileButton.Location = new System.Drawing.Point(410, 0);
            this.ChooseLocalFileButton.Name = "ChooseLocalFileButton";
            this.ChooseLocalFileButton.Size = new System.Drawing.Size(40, 31);
            this.ChooseLocalFileButton.TabIndex = 1;
            this.ChooseLocalFileButton.Text = "...";
            this.ChooseLocalFileButton.UseVisualStyleBackColor = true;
            this.ChooseLocalFileButton.Click += new System.EventHandler(this.ChooseLocalFileButton_Click);
            // 
            // LocalFileTextBox
            // 
            this.LocalFileTextBox.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.LocalFileTextBox.Location = new System.Drawing.Point(15, 5);
            this.LocalFileTextBox.Name = "LocalFileTextBox";
            this.LocalFileTextBox.Size = new System.Drawing.Size(389, 21);
            this.LocalFileTextBox.TabIndex = 0;
            // 
            // panel5
            // 
            this.panel5.Controls.Add(this.splitContainer5);
            this.panel5.Dock = System.Windows.Forms.DockStyle.Top;
            this.panel5.Location = new System.Drawing.Point(3, 17);
            this.panel5.Name = "panel5";
            this.panel5.Size = new System.Drawing.Size(452, 394);
            this.panel5.TabIndex = 2;
            // 
            // splitContainer5
            // 
            this.splitContainer5.Dock = System.Windows.Forms.DockStyle.Fill;
            this.splitContainer5.Location = new System.Drawing.Point(0, 0);
            this.splitContainer5.Name = "splitContainer5";
            // 
            // splitContainer5.Panel1
            // 
            this.splitContainer5.Panel1.Controls.Add(this.directoryTree);
            // 
            // splitContainer5.Panel2
            // 
            this.splitContainer5.Panel2.Controls.Add(this.filesList);
            this.splitContainer5.Size = new System.Drawing.Size(452, 394);
            this.splitContainer5.SplitterDistance = 182;
            this.splitContainer5.TabIndex = 0;
            // 
            // directoryTree
            // 
            this.directoryTree.Dock = System.Windows.Forms.DockStyle.Fill;
            this.directoryTree.ImageIndex = 0;
            this.directoryTree.ImageList = this.directoryIcons;
            this.directoryTree.Location = new System.Drawing.Point(0, 0);
            this.directoryTree.Name = "directoryTree";
            this.directoryTree.SelectedImageIndex = 0;
            this.directoryTree.Size = new System.Drawing.Size(182, 394);
            this.directoryTree.TabIndex = 0;
            this.directoryTree.BeforeExpand += new System.Windows.Forms.TreeViewCancelEventHandler(this.directoryTree_BeforeExpand);
            this.directoryTree.AfterExpand += new System.Windows.Forms.TreeViewEventHandler(this.directoryTree_AfterExpand);
            // 
            // directoryIcons
            // 
            this.directoryIcons.ImageStream = ((System.Windows.Forms.ImageListStreamer)(resources.GetObject("directoryIcons.ImageStream")));
            this.directoryIcons.TransparentColor = System.Drawing.Color.Transparent;
            this.directoryIcons.Images.SetKeyName(0, "Computer.ico");
            this.directoryIcons.Images.SetKeyName(1, "Closed Folder.ico");
            this.directoryIcons.Images.SetKeyName(2, "Open Folder.ico");
            this.directoryIcons.Images.SetKeyName(3, "Fixed Drive.ico");
            this.directoryIcons.Images.SetKeyName(4, "My Documents.ico");
            // 
            // filesList
            // 
            this.filesList.Activation = System.Windows.Forms.ItemActivation.OneClick;
            this.filesList.Dock = System.Windows.Forms.DockStyle.Fill;
            this.filesList.Location = new System.Drawing.Point(0, 0);
            this.filesList.Name = "filesList";
            this.filesList.Size = new System.Drawing.Size(266, 394);
            this.filesList.SmallImageList = this.filesIcons;
            this.filesList.TabIndex = 0;
            this.filesList.UseCompatibleStateImageBehavior = false;
            this.filesList.Click += new System.EventHandler(this.filesList_Click);
            // 
            // filesIcons
            // 
            this.filesIcons.ImageStream = ((System.Windows.Forms.ImageListStreamer)(resources.GetObject("filesIcons.ImageStream")));
            this.filesIcons.TransparentColor = System.Drawing.Color.Transparent;
            this.filesIcons.Images.SetKeyName(0, "Closed Folder.ico");
            this.filesIcons.Images.SetKeyName(1, "Computer.ico");
            this.filesIcons.Images.SetKeyName(2, "Fixed Drive.ico");
            this.filesIcons.Images.SetKeyName(3, "My Documents.ico");
            this.filesIcons.Images.SetKeyName(4, "Open Folder.ico");
            // 
            // RomateGroupBox
            // 
            this.RomateGroupBox.BackColor = System.Drawing.SystemColors.ControlLight;
            this.RomateGroupBox.Controls.Add(this.DownloadToLocalBtn);
            this.RomateGroupBox.Controls.Add(this.panel8);
            this.RomateGroupBox.Controls.Add(this.panel7);
            this.RomateGroupBox.Controls.Add(this.panel4);
            this.RomateGroupBox.Dock = System.Windows.Forms.DockStyle.Fill;
            this.RomateGroupBox.Location = new System.Drawing.Point(0, 0);
            this.RomateGroupBox.Name = "RomateGroupBox";
            this.RomateGroupBox.Size = new System.Drawing.Size(457, 495);
            this.RomateGroupBox.TabIndex = 0;
            this.RomateGroupBox.TabStop = false;
            this.RomateGroupBox.Text = "远程客户端";
            // 
            // DownloadToLocalBtn
            // 
            this.DownloadToLocalBtn.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.DownloadToLocalBtn.Font = new System.Drawing.Font("微软雅黑", 14.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(134)));
            this.DownloadToLocalBtn.Location = new System.Drawing.Point(155, 453);
            this.DownloadToLocalBtn.Name = "DownloadToLocalBtn";
            this.DownloadToLocalBtn.Size = new System.Drawing.Size(185, 39);
            this.DownloadToLocalBtn.TabIndex = 5;
            this.DownloadToLocalBtn.Text = "受控端->本地";
            this.DownloadToLocalBtn.UseVisualStyleBackColor = true;
            this.DownloadToLocalBtn.Click += new System.EventHandler(this.DownloadToLocalBtn_Click);
            // 
            // panel8
            // 
            this.panel8.BorderStyle = System.Windows.Forms.BorderStyle.FixedSingle;
            this.panel8.Controls.Add(this.RomateFileTextBox);
            this.panel8.Dock = System.Windows.Forms.DockStyle.Top;
            this.panel8.Location = new System.Drawing.Point(3, 411);
            this.panel8.Name = "panel8";
            this.panel8.Size = new System.Drawing.Size(451, 33);
            this.panel8.TabIndex = 2;
            // 
            // RomateFileTextBox
            // 
            this.RomateFileTextBox.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.RomateFileTextBox.Location = new System.Drawing.Point(3, 5);
            this.RomateFileTextBox.Name = "RomateFileTextBox";
            this.RomateFileTextBox.Size = new System.Drawing.Size(443, 21);
            this.RomateFileTextBox.TabIndex = 0;
            // 
            // panel7
            // 
            this.panel7.Dock = System.Windows.Forms.DockStyle.Top;
            this.panel7.Location = new System.Drawing.Point(3, 51);
            this.panel7.Name = "panel7";
            this.panel7.Size = new System.Drawing.Size(451, 360);
            this.panel7.TabIndex = 1;
            // 
            // panel4
            // 
            this.panel4.BorderStyle = System.Windows.Forms.BorderStyle.FixedSingle;
            this.panel4.Controls.Add(this.RomateUpdownGuidCB);
            this.panel4.Dock = System.Windows.Forms.DockStyle.Top;
            this.panel4.Location = new System.Drawing.Point(3, 17);
            this.panel4.Name = "panel4";
            this.panel4.Size = new System.Drawing.Size(451, 34);
            this.panel4.TabIndex = 0;
            // 
            // RomateUpdownGuidCB
            // 
            this.RomateUpdownGuidCB.Dock = System.Windows.Forms.DockStyle.Top;
            this.RomateUpdownGuidCB.Font = new System.Drawing.Font("Georgia", 14.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.RomateUpdownGuidCB.FormattingEnabled = true;
            this.RomateUpdownGuidCB.Location = new System.Drawing.Point(0, 0);
            this.RomateUpdownGuidCB.Name = "RomateUpdownGuidCB";
            this.RomateUpdownGuidCB.Size = new System.Drawing.Size(449, 31);
            this.RomateUpdownGuidCB.TabIndex = 0;
            // 
            // ClientDetailtabPage
            // 
            this.ClientDetailtabPage.Location = new System.Drawing.Point(4, 25);
            this.ClientDetailtabPage.Name = "ClientDetailtabPage";
            this.ClientDetailtabPage.Size = new System.Drawing.Size(925, 501);
            this.ClientDetailtabPage.TabIndex = 2;
            this.ClientDetailtabPage.Text = "详细信息";
            this.ClientDetailtabPage.UseVisualStyleBackColor = true;
            // 
            // OutputRichTextBox
            // 
            this.OutputRichTextBox.BackColor = System.Drawing.SystemColors.Info;
            this.OutputRichTextBox.Dock = System.Windows.Forms.DockStyle.Fill;
            this.OutputRichTextBox.Location = new System.Drawing.Point(0, 0);
            this.OutputRichTextBox.Name = "OutputRichTextBox";
            this.OutputRichTextBox.Size = new System.Drawing.Size(933, 188);
            this.OutputRichTextBox.TabIndex = 0;
            this.OutputRichTextBox.Text = "";
            // 
            // splitContainer3
            // 
            this.splitContainer3.Dock = System.Windows.Forms.DockStyle.Fill;
            this.splitContainer3.Location = new System.Drawing.Point(0, 0);
            this.splitContainer3.Name = "splitContainer3";
            this.splitContainer3.Orientation = System.Windows.Forms.Orientation.Horizontal;
            // 
            // splitContainer3.Panel1
            // 
            this.splitContainer3.Panel1.Controls.Add(this.panel3);
            // 
            // splitContainer3.Panel2
            // 
            this.splitContainer3.Panel2.Controls.Add(this.ClientStatus);
            this.splitContainer3.Size = new System.Drawing.Size(243, 730);
            this.splitContainer3.SplitterDistance = 533;
            this.splitContainer3.TabIndex = 0;
            // 
            // panel3
            // 
            this.panel3.Controls.Add(this.groupBox1);
            this.panel3.Dock = System.Windows.Forms.DockStyle.Fill;
            this.panel3.Location = new System.Drawing.Point(0, 0);
            this.panel3.Name = "panel3";
            this.panel3.Size = new System.Drawing.Size(243, 533);
            this.panel3.TabIndex = 0;
            // 
            // groupBox1
            // 
            this.groupBox1.BackColor = System.Drawing.SystemColors.Control;
            this.groupBox1.Controls.Add(this.ClientListView);
            this.groupBox1.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox1.Location = new System.Drawing.Point(0, 0);
            this.groupBox1.Name = "groupBox1";
            this.groupBox1.Size = new System.Drawing.Size(243, 533);
            this.groupBox1.TabIndex = 0;
            this.groupBox1.TabStop = false;
            this.groupBox1.Text = "客户端列表";
            // 
            // ClientListView
            // 
            this.ClientListView.Activation = System.Windows.Forms.ItemActivation.OneClick;
            this.ClientListView.Dock = System.Windows.Forms.DockStyle.Fill;
            this.ClientListView.Location = new System.Drawing.Point(3, 17);
            this.ClientListView.Name = "ClientListView";
            this.ClientListView.Size = new System.Drawing.Size(237, 513);
            this.ClientListView.TabIndex = 0;
            this.ClientListView.UseCompatibleStateImageBehavior = false;
            this.ClientListView.Click += new System.EventHandler(this.ClientListView_Click);
            // 
            // ClientStatus
            // 
            this.ClientStatus.BackColor = System.Drawing.SystemColors.ActiveCaption;
            this.ClientStatus.BorderStyle = System.Windows.Forms.BorderStyle.Fixed3D;
            this.ClientStatus.Dock = System.Windows.Forms.DockStyle.Fill;
            this.ClientStatus.Location = new System.Drawing.Point(0, 0);
            this.ClientStatus.Name = "ClientStatus";
            this.ClientStatus.Size = new System.Drawing.Size(243, 193);
            this.ClientStatus.TabIndex = 0;
            // 
            // splitContainer1
            // 
            this.splitContainer1.Dock = System.Windows.Forms.DockStyle.Fill;
            this.splitContainer1.FixedPanel = System.Windows.Forms.FixedPanel.Panel1;
            this.splitContainer1.Location = new System.Drawing.Point(0, 0);
            this.splitContainer1.Name = "splitContainer1";
            // 
            // splitContainer1.Panel1
            // 
            this.splitContainer1.Panel1.Controls.Add(this.splitContainer3);
            // 
            // splitContainer1.Panel2
            // 
            this.splitContainer1.Panel2.Controls.Add(this.splitContainer2);
            this.splitContainer1.Size = new System.Drawing.Size(1184, 730);
            this.splitContainer1.SplitterDistance = 243;
            this.splitContainer1.TabIndex = 0;
            // 
            // UpdateClientsInfoTimer
            // 
            this.UpdateClientsInfoTimer.Interval = 60000;
            this.UpdateClientsInfoTimer.Tick += new System.EventHandler(this.UpdateClientsInfoTimer_Tick);
            // 
            // UpdateCommandTimer
            // 
            this.UpdateCommandTimer.Interval = 1000;
            this.UpdateCommandTimer.Tick += new System.EventHandler(this.UpdateCommandTimer_Tick);
            // 
            // MainForm
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(1184, 730);
            this.Controls.Add(this.splitContainer1);
            this.Icon = ((System.Drawing.Icon)(resources.GetObject("$this.Icon")));
            this.Name = "MainForm";
            this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen;
            this.Text = "木马命令控制器";
            this.SizeChanged += new System.EventHandler(this.MainForm_SizeChanged);
            this.splitContainer2.Panel1.ResumeLayout(false);
            this.splitContainer2.Panel2.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.splitContainer2)).EndInit();
            this.splitContainer2.ResumeLayout(false);
            this.MainTabControl.ResumeLayout(false);
            this.CommandSendTabPage.ResumeLayout(false);
            this.panel1.ResumeLayout(false);
            this.groupBox2.ResumeLayout(false);
            this.panel2.ResumeLayout(false);
            this.UpDownloadTabPage.ResumeLayout(false);
            this.splitContainer4.Panel1.ResumeLayout(false);
            this.splitContainer4.Panel2.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.splitContainer4)).EndInit();
            this.splitContainer4.ResumeLayout(false);
            this.LocalGroupBox.ResumeLayout(false);
            this.panel6.ResumeLayout(false);
            this.panel6.PerformLayout();
            this.panel5.ResumeLayout(false);
            this.splitContainer5.Panel1.ResumeLayout(false);
            this.splitContainer5.Panel2.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.splitContainer5)).EndInit();
            this.splitContainer5.ResumeLayout(false);
            this.RomateGroupBox.ResumeLayout(false);
            this.panel8.ResumeLayout(false);
            this.panel8.PerformLayout();
            this.panel4.ResumeLayout(false);
            this.splitContainer3.Panel1.ResumeLayout(false);
            this.splitContainer3.Panel2.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.splitContainer3)).EndInit();
            this.splitContainer3.ResumeLayout(false);
            this.panel3.ResumeLayout(false);
            this.groupBox1.ResumeLayout(false);
            this.splitContainer1.Panel1.ResumeLayout(false);
            this.splitContainer1.Panel2.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.splitContainer1)).EndInit();
            this.splitContainer1.ResumeLayout(false);
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.SplitContainer splitContainer2;
        private System.Windows.Forms.TabControl MainTabControl;
        private System.Windows.Forms.TabPage CommandSendTabPage;
        private System.Windows.Forms.Panel panel2;
        private System.Windows.Forms.Panel AttrPanel;
        private System.Windows.Forms.Button SendCommandBtn;
        private System.Windows.Forms.Panel panel1;
        private System.Windows.Forms.GroupBox groupBox2;
        private System.Windows.Forms.ListView CommandListView;
        private System.Windows.Forms.TabPage UpDownloadTabPage;
        private System.Windows.Forms.RichTextBox OutputRichTextBox;
        private System.Windows.Forms.SplitContainer splitContainer3;
        private System.Windows.Forms.Panel panel3;
        private System.Windows.Forms.GroupBox groupBox1;
        private System.Windows.Forms.ListView ClientListView;
        private System.Windows.Forms.Panel ClientStatus;
        private System.Windows.Forms.SplitContainer splitContainer1;
        private System.Windows.Forms.TabPage ClientDetailtabPage;
        private System.Windows.Forms.Timer UpdateClientsInfoTimer;
        private System.Windows.Forms.SplitContainer splitContainer4;
        private System.Windows.Forms.GroupBox LocalGroupBox;
        private System.Windows.Forms.GroupBox RomateGroupBox;
        private System.Windows.Forms.Panel panel4;
        private System.Windows.Forms.ImageList directoryIcons;
        private System.Windows.Forms.ImageList filesIcons;
        private System.Windows.Forms.ComboBox RomateUpdownGuidCB;
        private System.Windows.Forms.Button UploadToClientBtn;
        private System.Windows.Forms.Panel panel6;
        private System.Windows.Forms.Button ChooseLocalFileButton;
        private System.Windows.Forms.TextBox LocalFileTextBox;
        private System.Windows.Forms.Panel panel5;
        private System.Windows.Forms.SplitContainer splitContainer5;
        private System.Windows.Forms.TreeView directoryTree;
        private System.Windows.Forms.ListView filesList;
        private System.Windows.Forms.Button DownloadToLocalBtn;
        private System.Windows.Forms.Panel panel8;
        private System.Windows.Forms.TextBox RomateFileTextBox;
        private System.Windows.Forms.Panel panel7;
        private System.Windows.Forms.Timer UpdateCommandTimer;
    }
}

