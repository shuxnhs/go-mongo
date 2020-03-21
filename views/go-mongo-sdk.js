// <script src="https://cdn.staticfile.org/jquery/1.10.2/jquery.min.js"></script>
// 必须先引入jquery才能使用ajax
const GO_MONGO_HOST = "http://127.0.0.1:8081";
const GO_MONGO_KEY = "1DD75A62EB5E561F0F10A9A51270E5A6";

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
    // let reqType = type.toUpperCase();
    // $.ajax({
    //     url: url,
    //     type: reqType,
    //     dataType: 'json',
    //     data: ""
    // })
    // .success(function(data) {
    //     apiData = JSON.stringify(data);
    //     return apiData
    //     alert(apiData)
    // });
    // return apiData
}
