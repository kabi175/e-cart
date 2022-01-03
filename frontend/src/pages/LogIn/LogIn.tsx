import { observer } from 'mobx-react-lite'
import React, { useState } from 'react'
import { Link } from 'react-router-dom'
import { useStoreContext } from '../../store/store'
function LogInPage() {
  const [email, setEmail] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const { authStore } = useStoreContext()
  return (
    <div className="container">
      <form
        className="field"
        onSubmit={(e) => {
          e.preventDefault()
        }}
      >
        <label className="label">
          {' '}
          User Name
          <input
            onChange={(event) => setEmail(event.target.value)}
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
            onChange={(event) => setPassword(event.target.value)}
            className="input"
            type="password"
            placeholder="********"
            required
          />
        </label>
        <div className="buttons">
          <button
            className="button  is-primary"
            onClick={() => authStore?.login(email, password)}
          >
            {' '}
            Login{' '}
          </button>
          <Link to="/signup" className="button is-outlined is-primary">
            {' '}
            Sign Up{' '}
          </Link>
        </div>
      </form>
    </div>
  )
}

export default observer(LogInPage)
