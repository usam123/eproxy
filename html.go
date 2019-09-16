package main

func html() string {
	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Title</title>
		
		<script src="https://cdn.jsdelivr.net/npm/vue@2.6.10/dist/vue.js"></script>
		<style>body{margin:0;background-color:#fafafa;font:14px 'Helvetica Neue',Helvetica,Arial,sans-serif}h2{margin:0;font-size:12px}a{color:#000;text-decoration:none}input{outline:0}button{margin:0;padding:0;border:0;background:0;font-size:100%;vertical-align:baseline;font-family:inherit;font-weight:inherit;color:inherit;outline:0}ul{padding:0;margin:0;list-style:none}.page-top{width:100%;height:40px;background-color:#db4c3f}.page-content{width:50%;margin:0 auto}.page-content h2{line-height:40px;font-size:18px;color:#fff}.main{width:50%;margin:0 auto;box-sizing:border-box}.task-input{width:99%;height:30px;outline:0;border:1px solid #ccc}.task-count{display:flex;margin:10px 0}.task-count li{padding-left:10px;flex:1;color:#dd4b39}.task-count li:nth-child(1){padding:5px 0 0 10px}.action{text-align:center;display:flex}.action a{margin:0 10px;flex:1;padding:5px 0;color:#777}.action a:nth-child(3){margin-right:0}.active{border:1px solid rgba(175,47,47,0.2)}.tasks{background-color:#fff}.no-task-tip{padding:10px 0 10px 10px;display:block;border-bottom:1px solid #ededed;color:#777}.big-title{color:#222}.todo-list{margin:0;padding:0;list-style:none}.todo-list li{position:relative;font-size:16px;border-bottom:1px solid #ededed}.todo-list li:hover{background-color:#fafafa}.todo-list li.editing{border-bottom:0;padding:0}.todo-list li.editing .edit{display:block;width:506px;padding:13px 17px 12px 17px;margin:0 0 0 43px}.todo-list li.editing .view{display:none}.todo-list li .toggle{text-align:center;width:40px;height:auto;position:absolute;top:5px;bottom:0;margin:auto 0;border:0;-webkit-appearance:none;appearance:none}.toggle{text-align:center;width:40px;height:auto;position:absolute;top:5px;bottom:0;margin:auto 0;border:0;-webkit-appearance:none;appearance:none}.toggle:after{content:url('data:image/svg+xml;utf8,<svgxmlns="http://www.w3.org/2000/svg"width="40"height="40"viewBox="-10-18100135"><circlecx="50"cy="50"r="40"fill="none"stroke="#ededed"stroke-width="3"/></svg>')}.toggle:checked:after{content:url('data:image/svg+xml;utf8,<svgxmlns="http://www.w3.org/2000/svg"width="40"height="40"viewBox="-10-18100135"><circlecx="50"cy="50"r="40"fill="none"stroke="#bddad5"stroke-width="3"/><pathfill="#5dc2af"d="M7225L42712756l-44202034-52z"/></svg>')}.todo-list li label{font-size: 90%;color: #6e6e6e;white-space:pre-line;word-break:break-all;padding:15px 60px 15px 15px;margin-left:45px;display:block;line-height:1.2;transition:color .4s}.todo-list li.completed label{color:#d9d9d9;text-decoration:line-through}.todo-list li .destroy{display:none;position:absolute;top:0;right:10px;bottom:0;width:40px;height:40px;margin:auto 0;font-size:30px;color:#cc9a9a;margin-bottom:11px;transition:color .2s ease-out}.todo-list li .destroy:hover{color:#af5b5e}.todo-list li .destroy:after{content:'×'}.todo-list li:hover .destroy{display:block}.todo-list li .edit{display:none}.todo-list li.editing:last-child{margin-bottom:-1px} .submit{cursor: pointer;text-decoration:none;background:#2f435e;color:#f2f2f2;width: 100%;padding: 10px 30px 10px 30px;font-size:16px;font-family: 微软雅黑,宋体,Arial,Helvetica,Verdana,sans-serif;font-weight:bold;border-radius:3px;-webkit-transition:all linear 0.30s;-moz-transition:all linear 0.30s;transition:all linear 0.30s;} .submit:hover{background:#385f9e;} </style>
	</head>
	<body>
<div class="page-top">
    <div class="page-content">
        <h2>服务器代理</h2>
    </div>
</div>
<div class="main">
    <h3 style="line-height: 0.1em;" class="big-title">源:</h3>
    <input
            placeholder="源地址"
            class="task-input"
            type="text"
            style="margin-bottom: 3px;"
            v-model="targetUrl"
    /><br><br>
    <h3 style="line-height: 0.1em;" class="big-title">代理:</h3>
    <input
            placeholder="代理地址"
            type="text"
            class="task-input"
            v-model="proxyUrl"
    /><br><br>

    <h3 style="line-height: 0.1em;" class="big-title">socks5代理：</h3>
    <input
            placeholder="socks5地址"
            class="task-input"
            type="text"
            style="margin-bottom: 3px;"
            v-model="socks5"
            
    /><br><br>

    <h3 style="line-height: 0.1em;" class="big-title">添加替换任务 <span style="font: 14px 'Helvetica Neue',Helvetica,Arial,sans-serif;color:#aaaaaa">(回车添加)</span>:</h3>
    <input
            placeholder="替换前字符"
            class="task-input"
            type="text"
            style="margin-bottom: 3px;"
            v-on:keyup.enter="enterFn"
            v-model="todo"
    /><br>
    <input
            placeholder="替换后字符"
            type="text"
            class="task-input"
            v-on:keyup.enter="enterFn"
            v-model="v"
    />

    <ul class="task-count">
        
        <!--
        <li>个替换对象</li>
        <li class="action">
            <a :class="{active:visibility!=='unCompleted'&&visibility!=='completed'}" href="#all">所有任务</a>
            <a :class="{active:visibility==='unCompleted'}" href="#unCompleted">未完成的任务</a>
            <a :class="{active:visibility==='completed'}" href="#completed">完成的任务</a>
        </li>-->
    </ul>

    

    <h3 class="big-title">替换列表 ({{unComplete}})：</h3>
    <div class="tasks">
 
        <span class="no-task-tip" v-show="!list.length">还没有添加任何任务</span>
        <ul class="todo-list" v-show="list.length">
            <li class="todo"
                v-for="item in filterData"
                v-bind:class="{completed:false,editing:item===edtorTodos}"
            >
                <div class="view">
                    <input class="toggle"
                           type="checkbox"
                           v-model="item.isComplete"
                    />
                    <label @dblclick="edtorTodo(item)">{{item.title}}  >>>  {{item.value}}</label>
                    <button
                            class="destroy"
                            @click="delFn(item)"
                    ></button>
                </div>
                <input
                        class="edit"
                        type="text"
                        v-focus="edtorTodos===item"
                        v-model="item.title"
                        @blur="edtoEnd(item)"
                        @keyup.enter="edtoEnd(item)"
                        @keyup.esc="cancelEdit(item)"
                />
            </li>
        </ul>
        <button class="submit" @click="submit()" >确认提交</button>
        <br/><br/>
        <button class="submit" @click="restart()" style="background: #db4c3f;">服务器重启</button>
        
        

    </div>
</div>
<script>
function checkUrl(urlString){
    if(urlString!=""){
        var reg=/(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?/;
        if(!reg.test(urlString)){
            return false
        } else {
            return true
        }
    }
}

function parseURL(url) { 
    var a =  document.createElement('a'); 
    a.href = url; 
    return { 
    source: url, 
    protocol: a.protocol.replace(':',''), 
    host: a.hostname, 
    port: a.port, 
    query: a.search, 
    params: (function(){ 
        var ret = {}, 
            seg = a.search.replace(/^\?/,'').split('&'), 
            len = seg.length, i = 0, s; 
        for (;i<len;i++) { 
            if (!seg[i]) { continue; } 
            s = seg[i].split('='); 
            ret[s[0]] = s[1]; 
        } 
        return ret; 
    })(), 
    file: (a.pathname.match(/\/([^\/?#]+)$/i) || [,''])[1], 
    hash: a.hash.replace('#',''), 
    path: a.pathname.replace(/^([^\/])/,'/$1'), 
    relative: (a.href.match(/tps?:\/\/[^\/]+(.+)/) || [,''])[1], 
    segments: a.pathname.replace(/^\//,'').split('/') 
    }; 
   }   
  
var store = {
    save(key,value){
        window.localStorage.setItem(key,JSON.stringify(value));
    },
    fetch(key){
     return JSON.parse(window.localStorage.getItem(key))||[];
    }
}

var vm = new Vue({
    el:".main",
    data:{
        list: store.fetch("storeData"),
        todo:'',
        v:'',
        targetUrl: store.fetch("targetUrl"),
        proxyUrl: store.fetch("proxyUrl"),
        socks5: store.fetch("socks5"),
        edtorTodos:'',
        beforeTitle:"",
        visibility:"all"
    },
    watch:{
        list:{
            handler:function(){
                store.save("storeData",this.list);
            },
            deep:true
        },
        targetUrl:{
            handler:function(){
                store.save("targetUrl",this.targetUrl);
            },
            deep:true
        },
        proxyUrl:{
            handler:function(){
                store.save("proxyUrl",this.proxyUrl);
            },
            deep:true
        },
        socks5:{
            handler:function(){
                store.save("socks5",this.socks5);
            },
            deep:true
        } 
    },
    methods:{
        enterFnDomain(ev){
            if(this.targetUrl=="" || this.proxyUrl==""){alert("源/代理都不能为空");return}
            console.log(this.targetUrl)
            console.log(this.proxyUrl)

            /* 
            if(d0.protocol != d.protocol) {
                
                this.list.push({
                    title:d0.protocol,
                    value:d.protocol,
                    isComplete:false
                });
            }
            */

           /* 
            this.list.push({
            title:(d0.host + (d0.port ? ':'+d0.port : '')),
            value:(d.host  + (d.port ? ':'+d.port : '')),
            isComplete:false
            });
            this.domain0 = "";
            this.domain = ""
            */
        },
        enterFn(ev){
            if(this.todo==""){return;}
            this.list.push({
                title:this.todo,
                value:this.v,
                isComplete:false
            });
            this.todo = "";
            this.v = ""
            
                
        },
        delFn(item){
            var index = this.list.indexOf(item);
            this.list.splice(index,1)
        },
        edtorTodo(item){
            this.beforeTitle = item.title;
            this.edtorTodos = item;
        },
        edtoEnd(item){
            this.edtorTodos="";
            // this.cancelEdit = this.edtorTodos;
        },
        cancelEdit(item){
            item.title = this.beforeTitle;
            this.beforeTitle = '';
            this.edtorTodos='';
        },
        ajax(url, data, fn) {
            var xhr = new XMLHttpRequest();
            that = this;
            xhr.open("POST", url, true);
           
            xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");  
            xhr.onreadystatechange = function() {
              if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 304)) {
                fn.call(this, xhr.responseText); 
              }
            };
            xhr.send(data);
          },
        submit() {
            if(!checkUrl(this.targetUrl)) {
                alert("源网址不完整");
                return
            }
            if(!checkUrl(this.proxyUrl)) {
                alert("代理网址不完整");
                return
            }

            this.ajax("/confHttp/", "json=" + JSON.stringify(this.list) + "&targetUrl=" + this.targetUrl + "&proxyUrl=" + this.proxyUrl + "&socks5=" + this.socks5, function(res){  // 路径要与服务器端一致，最后要加斜杠 不然会出现 post 301 到 get
                if(res=='ok') alert("提交完成"+res)
                else alert("提交失败")
            });
        },
        restart(){
            alert("已经发送重启指令，稍后请手动刷新")
            this.ajax("/restart/", "flag=restart", function(res){ 
                console.log(res)
            });
        }
    },
    directives:{
        "focus":{
            update(el,binding){
                if(binding.value){
                    el.focus();
                }
            }
        }
    },
    computed:{
        unComplete(){
        return  this.list.filter(item=>{
                return !item.isComplete
            }).length
        },
        filterData(){
            var filter = {
                all:function(list){
                    return list;
                },
                completed:function(list){
                    return list.filter(item=>{
                        return item.isComplete;
                    })
                },
                unCompleted:function(list){
                    return list.filter(item=>{
                        return !item.isComplete;
                    })
                }
            }
            return filter[this.visibility]?filter[this.visibility](this.list):this.list;
        }
 
    }
});
function hashFn(){
    var hash = window.location.hash.slice(1);
    vm.visibility = hash;
}
hashFn();
window.addEventListener('hashchange',hashFn);
</script>
</body>
</html>`

	return html
}
