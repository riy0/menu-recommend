import Vuex from 'vuex'

const appStore = () => {
    return new Vuex.Store({
        state: {
            menu_list: {},
            menu:{},
            ingredient_list: {}
        },
        mutations: {
            setMenuList(state, menus){
                state.menu_list = {...menus.data}
            }
            setMenu(state, menu){
                state.menu = menu.data
            },
            setIngredientList(state, ingredients){
                state.ingredient_list = {...ingredients.data}
            }
            clearMenu(state){
                state.menu = {}
            }
            clearIngredient(state) {
                state.ingredient_list = {}
            }
        }
    })
}

export default appStore
