<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="shortcut icon" href="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1536746227490&di=0e72724c83cdcb3123b23a0e108c6fe1&imgtype=0&src=http%3A%2F%2Fs14.sinaimg.cn%2Fmw690%2F002cmdUzzy7gHAVLMFn1d%26690" type="image/x-icon" />

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    .logo {
      background-image: url(https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1536680448409&di=105c73fb09bfc260e048dc89f02e90cb&imgtype=0&src=http%3A%2F%2Fcaiji.3g.cnfol.com%2Fwine%2Fimages%2F201808%2F31%2F1535680802121651.png);
      background-repeat: no-repeat;
      -webkit-background-size: 100px 100px;
      background-size: 100px 100px;
      background-position: center center;
      text-align: center;
      font-size: 42px;
      padding: 250px 0 70px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }

    header {
      padding: 100px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
  </style>
</head>

<body>
  <header>
    <h1 class="logo">Welcome to chaincodeUI</h1>
    <div class="description">
      chaincodeUI is a simple & powerful  web  which can  deploy and  update chaincode simply.
    </div>
    </header>
  <footer>
   {{range .Colletion}}  
       {{with .}}  
    <div class="author">
        <a href="http://{{.HerfURL}}">{{.HerfName}}</a>  
    </div>
      {{end}}  
       {{end}}  
  </footer>
  <div class="backdrop"></div>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
