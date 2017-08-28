<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>{{.PageTitle}}</title>
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <meta name="keywords" content="">
  <meta name="description" content="">
  <link rel="stylesheet" href="/static/layui/css/layui.css">
  <link rel="stylesheet" href="/static/css/global.css">
</head>
<body>

<div class="header">
  <div class="main">
    <a class="logo" href="/" title="Fly">Fly社区</a>
    <div class="nav">
      <a class="nav-this" href="/index">
        <i class="iconfont icon-wenda"></i>问答
      </a>
    </div>
    
    <div class="nav-user">      
     


      {{if .IsLogin}}
      <!-- 登入后的状态 -->
      <a class="avatar" href="user/index.html">
        <img src="http://tp4.sinaimg.cn/1345566427/180/5730976522/0">
        <cite>{{.UserInfo.Username}}</cite>
        <!-- <i>VIP2</i> -->
      </a> 
      <div class="nav">
        <a href="/user/set/"><i class="iconfont icon-shezhi"></i>设置</a>
        <a href="/logout/"><i class="iconfont icon-tuichu" style="top: 0; font-size: 22px;"></i>退了</a>
      </div>
      {{else}}
         <!-- 未登入状态 -->
          <a class="unlogin" href="/login"><i class="iconfont icon-touxiang"></i></a>
          <span><a href="/login">登入</a><a href="/reg">注册</a></span>
          <p class="out-login">
            <a href="" onclick="layer.msg('正在通过QQ登入', {icon:16, shade: 0.1, time:0})" class="iconfont icon-qq" title="QQ登入"></a>
            <a href="" onclick="layer.msg('正在通过微博登入', {icon:16, shade: 0.1, time:0})" class="iconfont icon-weibo" title="微博登入"></a>
          </p>   
      {{end}}
      
      <!-- <a class="avatar" href="../user/index.html">
        <img src="http://tp4.sinaimg.cn/1345566427/180/5730976522/0">
        <cite>贤心</cite>
        <i>VIP2</i>
      </a>
      <div class="nav">
        <a href="../user/set.html"><i class="iconfont icon-shezhi"></i>设置</a>
        <a href=""><i class="iconfont icon-tuichu" style="top: 0; font-size: 22px;"></i>退了</a>
      </div> -->
      
      
    </div>
  </div>
</div>
 {{.LayoutContent}}
<div class="footer">
  <p><a href="">XXX社区</a> 2017 &copy; <a href="#">XXXXXXX</a></p>
  <p>
    <a href="" target="_blank">XXXXX</a>
    <a href="" target="_blank">XXXXX</a>
    <a href="" target="_blank">XXXXX</a>
  </p>
</div>