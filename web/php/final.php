<?php
echo "<pre>";
echo "The following postback has been delivered:";
echo "<br>";
echo "</pre>";
echo pretty_print($_POST['payload']);

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
