rule ExampleRule : tag1 tag2 tag3
{
    meta:
        author = "John Doe"
        description = "This is an example rule"
    strings:
        $a = "example"
        $b = { 6A 40 68 00 30 }
    condition:
        $a or $b
}