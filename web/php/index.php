<?php

if(count($_POST) > 0)
{
    if($_POST['payload'] != '')
    {
        header('Location: final.php');
    }
    else
    {
		$payloadError = 'validation';
    }
}


?>
<!DOCTYPE html>
<html>
    <head>
        <title>Postback Form</title>
		<link href="assets/styles.css" rel="stylesheet" type="text/css" />
    </head>
    <body>
		<div id="Header">
			<h2>
				Postback Form
			</h2>
		</div>        
        <div id="Body">
            <form method="post" action="final.php" >	
                <div class="<?=$payloadError?>">
                    <label>Payload:</label>
                    <textarea name="payload"></textarea>
                </div>
                <div class="multiple">
                    <label>&nbsp;</label>
                    <input type="submit" name="submit" value="Deliver">
                </div>
            </form>
        </div>
	</body>
</html>