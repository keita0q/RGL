 $.ajaxSetup({
    timeout: 1000
});
 
 
 window.Api = {
    // ユーザー一覧情報の取得  
    ajaxGetUser : function(id){
        $.ajax({
            type: 'GET',
            url: './rest/v1/users/'+id,
            dataType: 'json'
        }).fail(function() {
            alert('error');
        }).done(function(res) {
            return res;
        });
    },
    //ユーザー登録
    ajaxPostUser : function(body){
        $.ajax({
            type: 'POST',
            url: './rest/v1/users',
            dataType: 'json',
            data: body
        }).fail(function() {
            alert('error');
        }).done(function(response) {
            $('#hoge').html(response);
        });
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
        });
    },
    
    //投稿取得（一件）
    ajaxGetSpare : function(id){
        $.ajax({
            type: 'GET',
            url: './rest/v1/spares/'+id,
            dataType: 'json'
        }).fail(function() {
            alert('error');
        }).done(function(response) {
            $('#hoge').html(response);
        });
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
        });
    },
    
    //投稿検索
    ajaxSearchSpares : function(body){
        $.ajax({
            type: 'GET',
            url: '/rest/v1/spares',
            dataType: 'json',
            data: body
        }).fail(function() {
            alert('error');
        }).done(function(response) {
            $('#hoge').html(response);
        });
    },
    
    //投稿保存
    ajaxSaveSpares : function(body){
        $.ajax({
            type: 'POST',
            url: './rest/v1/spares',
            dataType: 'json',
            data: body
        }).fail(function() {
            alert('error');
        }).done(function(response) {
            $('#hoge').html(response);
        });
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
        });
    }
 }