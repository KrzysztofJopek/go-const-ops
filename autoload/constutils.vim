"Get current function name
"Works only with vim-go plugin installed and for gofmt'ed code
"TODO Revrite it later using go ASC
function! constutils#GoGetFunctionName()
  let l:curr_pos = getpos('.')

  normal [[
  let l:function_name = matchlist(getline('.'), '\vfunc.* ([A-Za-z_-]+) ?\(')[1]

  call setpos('.', l:curr_pos)
  return l:function_name
endfunction
