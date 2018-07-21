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
// Create Time         :    2018/3/14 13:45:00
// Update Time         :    2018/3/14 13:45:00
// Class Version       :    v1.0.0.0
// Class Description   :    
// ===============================================================================
using System.Drawing;
// ===============================================================================
namespace TrojanCommandSender.Src
{
    public class CONFIG
    {
        public static string EXEDIR             = "Command";                // exe所在目录名
        public static int DEFAULT_MAX_WAITTIME  = 10 * 1000;                // 一个exe等待其输出的最大时间
        public static string CONFIG_FILE        = "config.json";            // 服务器配置文件
        public static string GENERATOR_EXE_FILE = "generate_cmd.exe";       // DB命令生成工具
        public static string COMMAND_TABLE_NAME = "command";                // 表名
        public static string CLIENTS_TABLE_NAME = "clients";                // 表名
        public static int UPDATE_CLIENT_TIMER   = 60 * 1000;
        public static int UPDATE_COMMAND_TIMER   = 1 * 1000;
        public static Color INFO_COLOR          = Color.Green;
        public static Color ERROR_COLOR         = Color.Red;
    }
}