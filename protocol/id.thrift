
namespace go oso.inf.ids
namespace java oso.inf.ids
namespace cpp oso.inf.ids
namespace php oso.inf.ids
namespace py oso.inf.ids

struct UUIDResponse{
    1: required ResponseCode code;
    2: required string msg;
    3: required list<string> result;
}

enum ResponseCode {
    EXCEPTION = -1,
    UNKNOWN = 0,
    OK = 1,
    ERROR = 2
}

service Generator{

    /*
        生成64位的流水ID,无规律不重复
        bizCode : 业务码
        token : 接入验证码
        num : 一次接口调用需要返回的ID个数
    */
    UUIDResponse getCurrentID(1: string bizCode, 2: string token, 3: i16 num);

    /*
        从小到大顺序生成32位ID号码。到达最大值后返回错误
        bizCode : 业务码
        token : 接入验证码
        max : 最大值
        num : 一次接口调用需要返回的ID个数
    */
    UUIDResponse getPrimaryID32(1: string bizCode, 2: string token, 3: i64 max, 4: i16 num);

    /*
        从大到小按照一定规律生成64位ID号码,到最小值后返回错误
        bizCode : 业务码
        token : 接入验证码
        max : 最大值
        num : 一次接口调用需要返回的ID个数
    */
    UUIDResponse getPrimaryID64(1: string bizCode, 2: string token, 3: i64 min, 4: i16 num);

    /*
        从1到limit顺序并且循环计数
        bizCode : 业务码
        token : 接入验证码
        limit : 最大次序为64位有符号数字最大值
        num : 一次接口调用需要返回的ID个数
    */
    UUIDResponse getOrder(1: string bizCode, 2: string token, 3: i64 limit, 4: i16 num);

}
