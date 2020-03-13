<?php
/**-------------------使用示例------------------------**/

// 引入文件
require_once dirname(__FILE__) . '/go-mongo-sdk.php';

// get请求
$getRs = GoMongoSDK::request('mongoR', 'CountData', array('collection' => 'test1', 'filter' => '{}'), 'GET');
print_r($getRs);

// post请求
$document = json_encode(array('name' => 'go-mongo'));
$postRs = GoMongoSDK::request('mongoC', 'CreateOneDocument', array('collection' => 'test1', 'document' => $document), 'POST');
print_r($postRs);

//Array
//(
//    [ret] => 200
//    [msg] =>
//    [data] => Array
//        (
//            [code] => 0
//            [msg] => 获取成功
//            [data] => 3
//        )
//)

//Array
//(
//    [ret] => 200
//    [msg] =>
//    [data] => Array
//        (
//            [code] => 0
//            [msg] => 新增成功
//            [data] => Array
//            (
//                [InsertedID] => 5e6b79ab1b4ca01269e93cc0
//            )
//
//        )
//)



