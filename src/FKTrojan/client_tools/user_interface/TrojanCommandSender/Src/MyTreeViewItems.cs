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
// Create Time         :    2018/3/16 14:36:56
// Update Time         :    2018/3/16 14:36:56
// Class Version       :    v1.0.0.0
// Class Description   :    
// ===============================================================================
using System;
using System.Diagnostics;
using System.IO;
using System.Windows.Forms;
// ===============================================================================
namespace TrojanCommandSender.Src
{
    public static class MyTreeViewItems
    {
        public static void Add(TreeNode e, ListView v)
        {
            //try..catch异常处理
            try
            {
                //判断"我的电脑"Tag 上面加载的该结点没指定其路径
                if (e.Tag.ToString() != "我的电脑")
                {
                    e.Nodes.Clear();                               //清除空节点再加载子节点
                    TreeNode tNode = e;                            //获取选中\展开\折叠结点
                    string path = tNode.Name;                      //路径  

                    //获取"我的文档"路径
                    if (e.Tag.ToString() == "我的文档")
                    {
                        path = Environment.GetFolderPath           //获取计算机我的文档文件夹
                            (Environment.SpecialFolder.MyDocuments);
                    }

                    //获取指定目录中的子目录名称并加载结点
                    string[] dics = Directory.GetDirectories(path);
                    foreach (string dic in dics)
                    {
                        TreeNode subNode = new TreeNode(new DirectoryInfo(dic).Name); //实例化
                        subNode.Name = new DirectoryInfo(dic).FullName;               //完整目录
                        subNode.Tag = subNode.Name;
                        subNode.ImageIndex = IconIndexes.ClosedFolder;       //设置获取节点显示图片
                        subNode.SelectedImageIndex = IconIndexes.OpenFolder; //设置选择节点显示图片
                        tNode.Nodes.Add(subNode);
                        subNode.Nodes.Add("");                               //加载空节点 实现+号
                    }

                    //ListView加载全部文件
                    String[] files = Directory.GetFiles(path);
                    ListViewItem item = null;
                    v.Items.Clear();
                    foreach (String file in files)
                    {
                        item = new ListViewItem(Path.GetFileName(file));
                        item.SubItems.Add(file);
                        v.Items.Add(item);
                    }
                }
            }
            catch (Exception msg)
            {
                MessageBox.Show(msg.Message);                   //异常处理
            }
        }
    }
}