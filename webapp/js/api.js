 $.ajaxSetup({
    timeout: 1000
});
 
 
 window.Api = {
    // ユーザー一覧情報の取得  
    ajaxGetUser : function(id,callback,comp){
        console.log("come");
        $.ajax({
            type: 'GET',
            url: 'rest/v1/users/'+id,
            dataType: 'json'
        }).done(callback).fail(function() {
            alert('error');
        }).complete(comp);
    },
    //ユーザー登録
    ajaxPostUser : function(body, callback){
        $.ajax({
            type: 'POST',
            url: 'rest/v1/users',
            dataType: 'json',
            data: JSON.stringify(body)
        }).fail(function() {
            alert('error');
        }).done(callback);
    },
    
    //ユーザー編集
    ajaxEditUser : function(body){
        $.ajax({
            type: 'PUT',
            url: './rest/v1/users',
            dataType: 'json',
            data: body
        }).fail(function() {
            alert('error');
        }).done(function(response) {
            $('#hoge').html(response);
        }).always(function(){
            alert('ok,ajax');
        });
    },
    
    //ユーザー削除
    ajaxDeleteUser : function(id){
        $.ajax({
            type: 'DELETE',
            url: '/rest/v1/users/'+id,
            dataType: 'json'
        }).fail(function() {
            alert('error');
        }).done(function(response) {
            $('#hoge').html(response);
        }).always(function(){
            alert('ok,ajax');
        });
    },
    
    //投稿取得（一件）
    ajaxGetSpare : function(id,callback){
        $.ajax({
            type: 'GET',
            url: './rest/v1/spares/'+id,
            dataType: 'json'
        }).fail(function() {
            alert('error');
        }).done(callback);
    },
    
    //投稿取得（一覧）
    ajaxGetSpares : function(userId){
        $.ajax({
            type: 'GET',
            url: '/rest/v1/spares/users/'+userId,
            dataType: 'json'
        }).fail(function() {
            alert('error');
        }).done(function(response) {
            $('#hoge').html(response);
        }).always(function(){
            alert('ok,ajax');
        });
    },
    
    //投稿検索
    ajaxSearchSpares : function(body, callback){
        $.ajax({
            type: 'GET',
            url: './rest/v1/spares',
            dataType: 'json',
            data: body
        }).fail(function() {
            alert('error');
        }).done(callback);
    },
    
    
    
    //投稿保存
    ajaxSaveSpares : function(body, callback){
        //console.log(JSON.stringify(body));
        $.ajax({
            type: 'POST',
            url: './rest/v1/spares',
            dataType: 'json',
            data: JSON.stringify(body)
        }).fail(function() {
            alert('error');
        }).done(callback);
    },
    
    //投稿編集
    ajaxEditSpares : function(body){
        $.ajax({
            type: 'PUT',
            url: './rest/v1/spares',
            dataType: 'json',
            data: body
        }).fail(function() {
            alert('error');
        }).done(function(response) {
            $('#hoge').html(response);
        }).always(function(){
            alert('ok,ajax');
        });
    },
    
    //投稿削除
    ajaxDeleteSpares : function(id){
        $.ajax({
            type: 'DELETE',
            url: '/rest/v1/spares/'+id,
            dataType: 'json'
        }).fail(function() {
            alert('error');
        }).done(function(response) {
            $('#hoge').html(response);
        }).always(function(){
            alert('ok,ajax');
        });
    }
 }