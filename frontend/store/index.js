import Vuex from 'vuex'

const appStore = () => {
    return new Vuex.Store({
        state: {
            menu_list: {},
            menu:{}
        },
        mutations: {
            setMenuList(state, menus){
                state.menu_list = {...menus.data}
            }
            setMenu(state, menu){
                state.menu = menu.data
            },
            clearMenu(state){
                state.menu = {}
            }
        }
    })
}

export default appStore
