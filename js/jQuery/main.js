// Check the main container is ready
		$('#feed').ready(function(){
			// Get each div
			$('.content').each(function(){
				// Get the content
				var str = $(this).html();
				// Set the regex string
				var regex = /(https?:\/\/([-\w\.]+)+(:\d+)?(\/([\w\/_\.]*(\?\S+)?)?)?)/g
				// Replace plain text links by hyperlinks
				var replaced_text = str.replace(regex, "<a href=$1 style=color:#66CC00>Available</a>");
				$(this).html(replaced_text);
			});
			
			$('.content').each(function(){
				// Get the content
				var str = $(this).html();
				// Set the regex string
				
				var regex = /(Not.Available)/g
				// Replace plain text links by hyperlinks
				var replaced_text = str.replace(regex, "<span style=color:red>$1</span>");
				$(this).html(replaced_text);
			});
		});
	(function(d, s, id) {
		  var js, fjs = d.getElementsByTagName(s)[0];
		  if (d.getElementById(id)) return;
		  js = d.createElement(s); js.id = id;
		  js.src = "//connect.facebook.net/en_GB/sdk.js#xfbml=1&appId=472036546284450&version=v2.3";
		  fjs.parentNode.insertBefore(js, fjs);
		}(document, 'script', 'facebook-jssdk'));