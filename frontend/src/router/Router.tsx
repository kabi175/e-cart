import React from 'react'
import { BrowserRouter, Switch, Route } from 'react-router-dom'
import SafeRoute from './SafeRoute'
import LogIn from '../pages/LogIn/LogIn'
import Home from '../pages/Home/Home'
import Product from '../pages/Product/Product'
import Cart from '../pages/Cart/Cart'
import Navbar from '../shared/Navbar/Navbar'
const Router: React.FC = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/">
          <SafeRoute>
            <Navbar />
            <Home />
          </SafeRoute>
        </Route>
        <Route exact path="/login">
          <LogIn />
        </Route>
        <Route path="/product">
          <Navbar />
          <Product />
        </Route>
        <Route exact path="/cart">
          <Navbar />
          <Cart />
        </Route>
      </Switch>
    </BrowserRouter>
  )
}

export default Router
