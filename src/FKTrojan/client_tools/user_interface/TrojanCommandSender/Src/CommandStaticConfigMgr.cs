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
using System.Collections.Generic;
using System.Linq;
using System.Web.Script.Serialization;
// ===============================================================================
namespace TrojanCommandSender.Src
{
    public class CommandStaticConfigMgr
    {
        public List<CommandStaticStruct> StaticCommandList { get; set; }
        public CommandStaticConfigMgr()
        {
            StaticCommandList = new List<CommandStaticStruct>();
            Clear();
        }
        public bool ParserCommandHelpOutput(string outputStr)
        {
            try
            {
                var jss = new JavaScriptSerializer();
                CommandStaticStruct s = jss.Deserialize<CommandStaticStruct>(outputStr);
                if (!s.IsValid())
                {
                    return false;
                }
                StaticCommandList.Add(s);
            }
            catch
            {
                return false;
            }
            return true;
        }
        public void SortCommandStaticList()
        {
            StaticCommandList = StaticCommandList.OrderBy(o => o.name).ToList();
        }
        public void Clear()
        {
            StaticCommandList.Clear();
        }
        public CommandStaticStruct GetCommandStaticStructByName(string strName)
        {
            CommandStaticStruct result = StaticCommandList.Find(x => x.name == strName);
            return result;
        }
    }
}