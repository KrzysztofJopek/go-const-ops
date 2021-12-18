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



let g:constutils#err_list = []

let g:constutils#wrap_prefix = 'ErrGo'
let g:constutils#wrap_function = 'wrapper.NewErrorWrap'
let g:constutils#err_package = 'wrapper'
function! constutils#WrapError()
  "Get function name and capitalize first letter
  let l:function_name = constutils#GoGetFunctionName()
  let l:function_name = toupper(l:function_name[0]) . l:function_name[1:]

  let l:default_err = printf("%s%s", g:constutils#wrap_prefix, l:function_name)
  let l:err_name = input('Enter error name', l:default_err)

  let l:command = printf("i%s(%s.%s, ea)", g:constutils#wrap_function, g:constutils#err_package, l:err_name)
  execute 'normal! ' . l:command
  call add(g:constutils#err_list, l:err_name)
endfunction

   
