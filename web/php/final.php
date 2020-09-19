<?php
error_reporting(-1);
ini_set('display_errors', 'On');

echo "<pre>";
echo "The postback has been delivered. Result: ";
echo "<br>";
echo call_post_back();
echo "<br>";
echo "Payload:";
echo "</pre>";
echo pretty_print($_POST['payload']);

function call_post_back()
{
    $url = "http://localhost:4000/postback";
    $curl = curl_init($url);
    curl_setopt($curl, CURLOPT_HTTPHEADER, array('Content-Type:application/json'));
    curl_setopt($curl, CURLOPT_POST, true);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($curl, CURLOPT_POSTFIELDS, $_POST['payload']);
    $response = curl_exec($curl);
    $status = curl_getinfo($curl, CURLINFO_HTTP_CODE);
    echo "Http status: {$status} - Http response: {$response}";
    curl_close($curl);
}
function pretty_print($json_data)
{

$space = 0;
$flag = false;

echo "<pre>";

for($counter=0; $counter<strlen($json_data); $counter++)
{

if ( $json_data[$counter] == '}' || $json_data[$counter] == ']' )
{
$space--;
echo "\n";
echo str_repeat(' ', ($space*2));
}

if ( $json_data[$counter] == '"' && ($json_data[$counter-1] == ',' ||
$json_data[$counter-2] == ',') )
{
echo "\n";
echo str_repeat(' ', ($space*2));
}
if ( $json_data[$counter] == '"' && !$flag )
{
if ( $json_data[$counter-1] == ':' || $json_data[$counter-2] == ':' )

echo '<span style="color:blue;font-weight:bold">';
else

echo '<span style="color:red;">';
}
echo $json_data[$counter];

if ( $json_data[$counter] == '"' && $flag )
echo '</span>';
if ( $json_data[$counter] == '"' )
$flag = !$flag;

if ( $json_data[$counter] == '{' || $json_data[$counter] == '[' )
{
$space++;
echo "\n";
echo str_repeat(' ', ($space*2));
}
}
echo "</pre>";
}
?>
