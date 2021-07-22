import * as ReactDOM from 'react-dom'
import { RootStateProvider } from './store/store'
import App from './App'

ReactDOM.render(
  <RootStateProvider>
    <App />
  </RootStateProvider>,
  document.getElementById('root')
)
