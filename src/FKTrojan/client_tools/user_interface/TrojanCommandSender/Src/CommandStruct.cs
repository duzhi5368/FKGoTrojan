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
// Create Time         :    2018/3/13 11:58:51
// Update Time         :    2018/3/13 11:58:51
// Class Version       :    v1.0.0.0
// Class Description   :    
// ===============================================================================
using System;
using System.Collections.Generic;
// ===============================================================================
namespace TrojanCommandSender.Src
{
    public class CommandStaticAttrStruct
    {
        public string long_fmt { get; set; }
        public string short_fmt { get; set; }
        public string example { get; set; }
        public string desc { get; set; }
        public bool required { get; set; }
        public string type { get; set; }

        public CommandStaticAttrStruct()
        {
            long_fmt = short_fmt = example = desc = type = string.Empty;
            required = false;
        }
    }
    public class CommandStaticStruct
    {
        public string name { get; set; }
        public string version { get; set; }
        public string desc { get; set; }
        public List<CommandStaticAttrStruct> Parameters { get; set; }

        public CommandStaticStruct()
        {
            Parameters = new List<CommandStaticAttrStruct>();
            Clear();
        }
        public void Clear()
        {
            name = version = desc = string.Empty;
            Parameters.Clear();
        }
        public bool IsValid()
        {
            if (string.IsNullOrEmpty(name))
                return false;
            return true;
        }
    }

    public class CommandDynamicAttrStruct
    {
        public string Value { get; set; }
        public Type TheType { get; set; }

        public CommandDynamicAttrStruct()
        {
            Value = string.Empty;
        }

        public T CastValue<T>()
        {
            return (T)Convert.ChangeType(Value, typeof(T));
        }
    }

    public class CommandDynamicStruct
    {
        public List<CommandDynamicAttrStruct> Parameters { get; set; }

        public CommandDynamicStruct()
        {
            Parameters = new List<CommandDynamicAttrStruct>();
            Clear();
        }
        public void Clear()
        {
            Parameters.Clear();
        }
    }
}