$( document ).ready(function() {
 
	$("#bills").keyup(function(){ 
	   //alert("Hello");
	   if(checkEmpty()){
			$("#pay").val(computeValuePerPax());
	   }
	});
});

function computeValuePerPax(){
	
	var payment = $("#bills").val();
	//discount
	payment = payment * (1-($("#discount").val() /100));
	//service charge
	payment = payment * (1+($("#serviceCharge").val()/100));
	//gst
	payment = payment * (1+($("#gst").val()/100));
	
	payment = payment / $("#numOfPpl").val();
	return payment.toFixed(2) ;
}

function checkEmpty() {
	if ($("#bills").val().length > 0 && 
		$("#numOfPpl").val().length > 0 &&
		$("#discount").val().length > 0 &&
		$("#serviceCharge").val().length > 0 &&
		$("#gst").val().length > 0) {
		return true;
	}else{
		return false;
	}
}