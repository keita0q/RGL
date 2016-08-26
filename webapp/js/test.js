var user = jQuery.parseJSON(localStorage.getItem("loginUser"))
var tags={};
var i = 0;
$(function(){
    
    $('.a').hover(function(){
        $(this).find('.b').css('color','red');
        },function(){
           $(this).find('.b').css('color','black'); 
        });
        
    $('#login-show').click(function(){
        $('#comfirm-modal').fadeOut();
        $('#complete-modal').fadeIn();
        });
        
    $('#form').submit(function(){
        var tag=$('#tag-in').val();
        console.log();
        tags[i] = tag;
        
        $("<span class=\"close-button\" id=" + i + ">" + tag + "</span>").appendTo('.tags').click(function(a){
            //console.log(a.currentTarget.id);
            var index=$(this).attr("id");
            delete tags[index];
            console.log(index);
            console.log(tags);
            $('#'+index).remove();
        });
        i++;
        // Object.values(tags);
        return false;
    });
    
    $('#tag_add').submit(function(){
        var tag=$('#tags').val();
        console.log();
        tags[i] = tag;
        
        $("<span class=\"close-button\" id=" + i + ">" + tag + "</span>").appendTo('.tags').click(function(a){
            //console.log(a.currentTarget.id);
            var index=$(this).attr("id");
            delete tags[index];
            console.log(index);
            console.log(tags);
            $('#'+index).remove();
        });
        i++;
        // Object.values(tags);
        return false;
    });

    
});