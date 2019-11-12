$(document).ready(function(){
var session = window.sessionStorage;
var cookie_phone = getCookie('cookie_phone');
    var cookie_passwd = getCookie('cookie_passwd');
	if(cookie_phone!=null&&cookie_passwd!=null){
        var params = {};
        params['phone'] = cookie_phone;
        params['passwd'] = cookie_passwd;
        params['action'] = 'login_cert';
        $.ajax({
           type: 'post',
           url: '/logcontrol.php',
           data: params,
           async: false,
	success:function(data){
                var data = eval("(" + data + ")");
                if(data['userid'] == 0){
			if(data['stats'] === '0'){
                        $("#error").text("Sie haben sich abgemeldet");
                    }else{
                        $("#error").text("Benutzername oder Passwort ungültig. Bitte versuchen sie es erneut.");
                    }
                    $("#error").addClass("show");
                    setCookie("cookie_phone",-1);
                    return false;
                }else{
                    session.setItem('login_status',1);
		 session.setItem('userid',data['userid']);
		 session.setItem('userphone',cookie_phone);
		}
           },
           error: function(errormessage) {
                console.log('<?=_("�~L�~A�~T��~H��~P~M�~F�| ~A�~X��~P�正确�~G��~T~Y")?>');
                return false;
            }
        });
    }
	var coin_userid = session.getItem("userid");
	if (session.getItem("login_status") == "1") {
		var coin_data = {};
		coin_data['userid'] = coin_userid;
		$.ajax({
			type: 'post',
			url: '/get_coin.php',
			data: coin_data,
			async: false,
			success: function (coin_callback) {
				$("#coinnum_index").text("x " + coin_callback);
			}
		});
	}
})


function getCookie(name) {
	var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");

	if (arr = document.cookie.match(reg))

		return unescape(arr[2]);
	else
		return null;
}

function getPage(game, page, num) {//page页num个game信息
	var start = (page - 1) * num;
	var end = page * num;
	var html_games = [];
	game = game.slice(start, end);
	$.each(game, function (i, c) {
		var game_name = "'" + c.title + "'"
		html_games.push('<div class=""><div class="row game-box">');
		html_games.push('<div class="col-xs-6"><h4>' + c.title + '</h4></div>');
		html_games.push('<div class="col-xs-6 text-right"><h4>Preis: <b>' + c.price + '</b> Credit</h4></div>');

		html_games.push('<div class="col-xs-12 margin-bottom"><a onclick="sendga(' + game_name + ')" href="details01.html?' + c.gameid + '"><img class="game-img" alt="' + c.title + '" src="img/' + c.pic_url + '/400X244.jpg"></a></div>');


		html_games.push('<div class="col-xs-12"><a onclick="sendga(' + game_name + ')"href="details01.html?' + c.gameid + '" class="button stretch">Wählen</a></div></div></div>');
	});
	$("#page").html(html_games.join(""));
};
function setDis(page, n) {//判断左右箭头是否可点击
	if (page == 1) {
		$("#prePage").addClass("disabled");
	} else {
		$("#prePage").removeClass("disabled");
	}
	if (page == n) {
		$("#nextPage").addClass("disabled");
	} else {
		$("#nextPage").removeClass("disabled");
	}
};


$(document).ready(function () {
	var session = window.sessionStorage;
	var userid = session.getItem("userid");
	var userphone = session.getItem("userphone");
	if (userid == null || userid == "") {
		userid = "null"
	}
	var num = 5;//每页显示条数
	var pagedata = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
	var pages = [];
	var params = {};
	params['action'] = 'get_list';
	var game;
	$.ajax({
		type: 'post',
		url: '/logcontrol.php',
		data: params,
		async: false,
		success: function (data) {
			console.log(data);
			game = eval("(" + data + ")");
		}
	});
	getPage(game, 1, num);
	var len = game.length;
	var n = Math.ceil(len / num);
	setDis(1, n);
	pagedata = pagedata.slice(0, n);
	$.each(pagedata, function (i, content) {
		pages.push('<option>' + content + '</option>');
	});
	$("#page_select").html(pages.join(""));

	$("#page_select").change(function () {
		var page = $(this).val();
		//ga('send', 'event', userid, "turnpage", page, 1);
		ga('send', 'event', userphone, "turnpage", page, 1);
		getPage(game, page, num);
		setDis(page, n);
	});
	$("#prePage").click(function () {
		if ($(this).attr("class") !== 'disabled') {
			var page = $("#page_select").val() - 1;
			//ga('send', 'event', userid, "turnpage", page, 1);
			ga('send', 'event', userphone, "turnpage", page, 1);
			getPage(game, page, num);
			$("#page_select").val(page);
			setDis(page, n);
		}
	});
	$("#nextPage").click(function () {
		if ($(this).attr("class") !== 'disabled') {
			var page = parseInt($("#page_select").val()) + 1;
			//ga('send', 'event', userid, "turnpage", page, 1);
			ga('send', 'event', userphone, "turnpage", page, 1);
			getPage(game, page, num);
			$("#page_select").val(page);
			setDis(page, n);
		}
	});
	$(".person").click(function () {
		var session = window.sessionStorage;
		//ga('send', 'event', userid, "tomygame", "tomygame", 1);
		ga('send', 'event', userphone, "tomygame", "tomygame", 1);
		if (session.getItem('login_status') == 1) {
			window.location = 'person.html';
		} else {
			session.setItem('his_url', 'person.html');
			window.location = 'login.html';
		}
	});
	/*getpage(1);*/
});

// $(document).ready(function(){
// 	getpage(1);
// }).on("a", "tap", function(e){
// 	e.preventDefault();
// 	var page = $(this).attr("data-value");
// 	getpage(page);
// });
function sendga(id) {
	var session = window.sessionStorage;
	var userid = session.getItem("userid");
	var userphone = session.getItem("userphone");
	// if (userid == null||userid==""){
	// 	userid = "null"
	// }
	// ga('send', 'event', userid, "play_index", id, 1);
	if (userphone == null || userphone == "") {
		userphone = "null"
	}
	ga('send', 'event', userphone, "play_index", id, 1);
}
