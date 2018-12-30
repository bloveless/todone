import * as React from 'react';
import * as ReactDOM from 'react-dom';
import CssBaseline from '@material-ui/core/CssBaseline';
import './styles/app.scss';

import Home from './pages/Home';

ReactDOM.render(
    <React.Fragment>
        <CssBaseline />
        <Home />
    </React.Fragment>,
    document.getElementById('react-root')
);
