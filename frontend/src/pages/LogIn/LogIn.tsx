import React from 'react'

function LogInPage() {
  return (
    <>
      <form className="field">
        <label className="label">
          {' '}
          User Name
          <input
            className="input"
            type="text"
            placeholder="e.g.  bob@gmain.com"
            required
          />
        </label>{' '}
        <label className="label">
          {' '}
          Password
          <input
            className="input"
            type="password"
            placeholder="********"
            required
          />
        </label>
        <div className="buttons">
          <button className="button  is-primary"> Login </button>
          <button className="button is-outlined is-primary"> Sign Up </button>
        </div>
      </form>
      <button className="button is-primary">
        <span className="icon is-small">
          <i className="fas fa-italic"></i>
        </span>
      </button>
    </>
  )
}

export default LogInPage
