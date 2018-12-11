import React, {Component} from 'react';
import './App.less';
import {BrowserRouter as Router, Route} from 'react-router-dom';
import {Provider, observer} from 'mobx-react';
import stores from './stores';
import Group from './page/Group'


@observer
class App extends Component {
    render() {
        return (
                <Provider stores={stores}>
                    <Router>
                        <div className="App">
                        <Route path='/' component={Group}/>
                        </div>
                    </Router>
                </Provider>
        );
    }
}

export default App;
