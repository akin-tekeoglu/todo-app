$( document ).ready(function() {
    $(".card").click(function() {
        $(".btn-group").hide();
        $(".card").removeClass("bg-info");
        $(this).find(".btn-group").show();
        $(this).addClass("bg-info");
    });
 });