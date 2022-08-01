<?php
function tab_expand($text)
{
    while (strstr($text, "\t")) {
        $text = preg_replace_callback('/^([^\t\n]*)(\t+)/m', 'tab_expand_helper', $text);
    }
    return $text;
}

function tab_expand_helper($matches)
{
    $tab_stop = 8;

    return $matches[1] . str_repeat(' ', strlen($matches[2]) * $tab_stop - (strlen($matches[1]) % $tab_stop));
}

function tab_unexpand($text)
{
    $tab_stop = 8;
    $lines = explode("\n", $text);
    foreach ($lines as $i => $line) {
        // 将制表符扩展位空格
        $line = tab_expand($line);
        $chunks = str_split($line, $tab_stop);
        $chunkCount = count($chunks);
        // 扫描除最后一个字符块以外的所有其他字符块
        for ($j = 0; $j < $chunkCount - 1; $j++) {
            echo $chunks[$j] . "\n";
            $chunks[$j] = preg_replace('/ {2,}$/', "\t", $chunks[$j]);
            echo $chunks[$j] . "\n";
        }
        // 如果最后一个字符块是相当于制表符的空格
        // 将它转换位一个制表符；否则，保持不变
        if ($chunks[$chunkCount - 1] == str_repeat(' ', $tab_stop)) {
            $chunks[$chunkCount - 1] = "\t";
        }
        // 重新组合字符块
        $lines[$i] = implode('', $chunks);
    }
    // 重新组合文本行
    return implode("\n", $lines);
}

$msg = '白日依山尽  黄河入海流	欲穷千里目	更上一层楼 白日依山尽  黄/河入海流    /    欲穷千里目 更上一层楼';
// $spaced = tab_expand($msg);
$tabbed = tab_unexpand($msg);

// echo $spaced . "\n";
echo $tabbed . "\n";