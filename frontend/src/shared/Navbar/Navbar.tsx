import React, { useContext } from 'react'
import { StoreContext } from '../../store/store'
import { Link } from 'react-router-dom'

const Navbar: React.FC = ({}) => {
  const { authStore } = useContext(StoreContext)
  return (
    <nav
      className="navbar has-shadow  is-link"
      role="navigation"
      aria-label="main navigation"
    >
      <div className="container">
        <div className="navbar-brand">
          <p className="navbar-item has-text-weight-bold is-size-4">SaleCart</p>
        </div>

        {/* navbar start elements */}
        <div className="navbar-start">
          <div className="navbar-item">
            <div className="field is-grouped ">
              <p className="control  has-icons-right">
                <input className="input" type="text" />
                <span className="icon is-right">
                  <i className="fas fa-search"></i>
                </span>
              </p>
            </div>
          </div>
        </div>

        {/*  navbar-end elements */}
        <div className="navbar-end mr-6">
          {/*  login button */}
          {authStore?.userName == null ? (
            <div className="navbar-item">
              <div className="field is-grouped">
                <button className="button is-white has-text-weight-semibold has-text-link">
                  {' '}
                  Login
                </button>
              </div>
            </div>
          ) : (
            <></>
          )}
          {/* more options - navbar dropdown  */}
          <div className="navbar-item control is-grouped has-dropdown is-hoverable">
            <span className="navbar-link">more</span>
            <div className="navbar-dropdown is-boxed">
              <span className="navbar-item">Sell on SaleCart</span>
              <span className="navbar-item">Help</span>
            </div>
          </div>

          {/* cart icon  */}
          <div className="navbar-item">
            <Link to="/cart">
              <span className="field navbar-link is-arrowless is-grouped icon-text">
                <span className="is-large pr-2">
                  <i className="fas fa-shopping-cart"></i>
                </span>
                <span className="has-text-weight-semibold">Cart</span>
              </span>
            </Link>
          </div>
        </div>
      </div>
    </nav>
  )
}

export default Navbar
