
function ch_op(o,v){o.style.opacity=v;}

function changeColor(){
	// 获取所有行
	var trs=document.getElementById("tbl").getElementsByTagName("tr");
	// 为偶数行添加class属性，且不包括标题行
	for(var i=1;i<trs.length;i+=2){
		trs[i].className="even";
	}
}
