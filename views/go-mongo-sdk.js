// <script src="https://cdn.staticfile.org/jquery/1.10.2/jquery.min.js"></script>
// 必须先引入jquery才能使用ajax
const GO_MONGO_HOST = "http://127.0.0.1:8081";
const GO_MONGO_KEY = "674D5122FEBC0F030C2AD55C9ED25B77";

function request(router, service, params, type){
    let url = GO_MONGO_HOST + "/" + router.trim() + "/" + service.trim();
    params.set("mongoKey", GO_MONGO_KEY);
    return doRequest(url, type, params);
}

function doRequest(url, type, params){
    url += "?";
    params.forEach(function(value, key){
        url += key + "=" + value + "&"
    });
    return url;
}

function delCookie(key) {
    var date = new Date();
    date.setTime(date.getTime() - 1);
    var delValue = getCookie(key);
    if (!!delValue) {
        document.cookie = key+'='+delValue+';expires='+date.toGMTString();
    }
}

function getCookie(key) {
    var arr,reg = RegExp('(^| )'+key+'=([^;]+)(;|$)');
    if (arr = document.cookie.match(reg))
        return decodeURIComponent(arr[2]);
    else
        return null;
}

