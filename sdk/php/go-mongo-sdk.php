<?php

defined('GO_MONGO_HOST') || define('GO_MONGO_HOST', 'TODO');    // TODO: 项目地址
defined('GO_MONGO_VERSION') || define('GO_MONGO_VERSION', 'v1');
defined('GO_MONGO_KEY') || define('GO_MONGO_KEY', 'TODO');      // TODO：填写项目的mongo-key

class GoMongoSDK{

    /**
     * @function: 发起go-mongo请求
     * @param   string  $router     请求的模块：mongoC，mongoU，mongoR，mongoD，mongodb，project
     * @param   string  $service    请求的接口名称
     * @param   array   $params     传递的参数数组
     * @param   string  $type       发送请求类型：GET/POST
     * @return  array
     */
    public static function request(string $router, string $service, array $params, string $type){
        $url = trim(GO_MONGO_HOST. '/') . trim(GO_MONGO_VERSION.'/') . $router. '/'. $service;
        $params['mongoKey'] = GO_MONGO_KEY;
        $rs = self::doRequest($url, $type, $params);
        return json_decode($rs, true);
    }


    /**
     * @function    发送get/post请求
     * @param  string    $url           请求地址
     * @param  string    $type          GET/POST
     * @param  array     $data          传递数据
     * @param  array     $header        请求头
     * @param  int       $ignoreSsl     https
     * @param  int       $timeout       过期时间
     * @return string/array
     */
    protected static function doRequest($url, $type = 'GET', $data=array(), $header=array(), $ignoreSsl=0, $timeout=30){
        $curl = curl_init();
        if(strtoupper($type) == 'GET'){
            if(!empty($data)){
                $url .= "?";
                foreach ($data as $k => $v){
                    $url .= "$k=$v&";
                }
            }
            // request url
            curl_setopt($curl, CURLOPT_URL, $url);
        }else{
            curl_setopt($curl, CURLOPT_URL, $url);
            curl_setopt($curl, CURLOPT_POST, 1);
            curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
        }
        // headers
        if(!empty($header)){
            curl_setopt($curl, CURLOPT_HTTPHEADER, $header);
        }
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($curl, CURLOPT_FOLLOWLOCATION, 1);

        // https
        if (!empty($ignoreSsl)) {
            curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, FALSE);
            curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, FALSE);
        }
        // timeout
        curl_setopt($curl, CURLOPT_CONNECTTIMEOUT, $timeout);
        curl_setopt($curl, CURLOPT_TIMEOUT, $timeout);

        $output = curl_exec($curl);
        if (curl_error($curl)){
            $output = curl_error($curl);
        }
        curl_close($curl);
        return $output;
    }

}
