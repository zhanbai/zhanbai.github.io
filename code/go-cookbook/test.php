<?php
$books = [
    ['Elmer Gantry', 'Sinclair Lewis', 1927],
    ["The Scarlatti Inheritance", "Robert Ludlum", "1971"],
    ["The Parsifal Mosaic", "William Styron", "1979"],
];

foreach ($books as $book) {
    // print pack('A25A15A4', $book[0], $book[1], $book[2]) . "\n";
    $title = str_pad(substr($book[0], 0, 25), 25, '.');
    $author = str_pad(substr($book[1], 0, 15), 15, '.');
    $year = str_pad(substr($book[2], 0, 4), 4, '.');
    print "$title$author$year\n";
}