# https://leetcode.com/discuss/74129/the-longest-absolute-path-in-file-system

def count_leading_tab(input, start):
    ret = 0
    while start < len(input) and input[start] == '\t':
        start += 1
        ret += 1
    return ret

def check_line(input, start):
    length = 0
    is_file = False
    while start < len(input) and input[start] != '\n':
        length += 1
        if input[start] == '.':
            is_file = True
        start += 1
    return length, is_file

def longest_in_one_path(input, start, leading, tabing):
    ret = 0
    while start < len(input):
        leading_tab = count_leading_tab(input, start)
        if leading_tab != tabing:
            break
        start = start+tabing
        length, is_file = check_line(input, start)
        start = length + start + 1
        if is_file:
            if leading+length > ret:
                ret = leading+length
        else:
            r, start = longest_in_one_path(input, start, leading+length+1, tabing+1)
            if r > ret:
                ret = r
    return ret, start

def longest_path(input):
    ret, _ = longest_in_one_path(input, 0, 0, 0)
    return ret

print longest_path("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext")
