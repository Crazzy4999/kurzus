<script setup lang="ts">
import CartImg from "@/assets/icons/8679184_shopping_basket_shop_icon.svg"
import MenuImg from "@/assets/icons/134216_menu_lines_hamburger_icon.svg"
import Sidebar from "@/components/Sidebar.vue"
import { ref } from "vue";

let toggleSidebar = ref(false)
</script>

<template>
    <header>
        <RouterLink to="/cart" class="cart-wrapper-mobile">
            <!--Icon by RmixIcon - https://www.iconfinder.com/remix-icon-->
            <img :src="CartImg" alt="cart">
        </RouterLink>
        <RouterLink to="/" class="logo-wrapper">
            <h3>Hangry</h3>
        </RouterLink>
        <nav>
            <ul>
                <li>
                    <RouterLink to="/login" class="nav-item">Login</RouterLink>
                </li>
                <li>
                    <RouterLink to="/signup" class="nav-item">Sign Up</RouterLink>
                </li>
                <li>
                    <RouterLink to="/suppliers" class="nav-item">Suppliers</RouterLink>
                </li>
                <li>
                    <RouterLink to="/categories" class="nav-item">Categories</RouterLink>
                </li>
                <li>
                    <RouterLink to="/cart" class="cart-wrapper">
                        <!--Icon by RmixIcon - https://www.iconfinder.com/remix-icon-->
                        <img :src="CartImg" alt="cart">
                    </RouterLink>
                </li>
            </ul>
        </nav>
        <button @click="toggleSidebar = !toggleSidebar">
            <!--Icon by Thimoty Miller - https://www.iconfinder.com/tmthymllr-->
            <img class="menu-img" :src="MenuImg" alt="menu">
        </button>
    </header>
    <Sidebar :toggle="toggleSidebar"/>
</template>

<style scoped>
@import "@/assets/layout.css";

header {
    z-index: 2;
    position: sticky;
    display: flex;
    justify-content: space-between;
    border-bottom-left-radius: var(--sub-p-size);
    border-bottom-right-radius: var(--sub-p-size);
    background-color: var(--second-color);
    box-shadow: 0 0 var(--sub-p-size) calc(var(--border-size) * 2) var(--settings-color);
    height: var(--h1-size);
    top: 0;
    padding-inline: var(--p-size);
}

.cart-wrapper {
    display: grid;
}

.cart-wrapper-mobile {
    display: none;
}

img {
    width: var(--h2-size);
    height: 70%;
    margin: auto;
}

.logo-wrapper {
    display: grid;
}

h3 {
    font-size: var(--h2-size);
    color: var(--main-color);
    margin-block: auto;
}

button {
    z-index: 4;
    display: none;
}

@media only screen and (hover: none) and (pointer: coarse) {
    header {
        z-index: unset;
        position: relative;
        height: fit-content;
    }

    .logo-wrapper {
        margin: var(--h4-size);
    }

    nav {
        display: none;
    }

    .cart-wrapper-mobile {
        display: grid;
    }

    button {
       display: inline-block;
    }
}
</style>