using System;
using System.Collections.Generic;
using System.Linq;
using System.Windows.Forms;

namespace TrojanCommandSender
{
    static class Program
    {
        public static MainForm mainForm;
        /// <summary>
        /// 应用程序的主入口点。
        /// </summary>
        [STAThread]
        static void Main()
        {
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            mainForm = new MainForm();
            mainForm.Init();
            Application.Run(mainForm);
        }
    }
}
