const menuClose = document.getElementById('menu_close')
const menuIconBtnOpen = document.getElementById('menu_icon_btn_open')
const menuOpen = document.getElementById('menu_open')
const menuIconBtnClose = document.getElementById('menu_icon_btn_close')

menuIconBtnOpen.addEventListener('click', () => {
  menuOpen.classList.add('menu_show')
  menuClose.classList.remove('menu_show')
})

menuIconBtnClose.addEventListener('click', () => {
  menuOpen.classList.remove('menu_show')
  menuClose.classList.add('menu_show')
})
