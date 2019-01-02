import * as React from 'react';
import * as ReactDOM from 'react-dom';
import CssBaseline from '@material-ui/core/CssBaseline';
import './styles/app.scss';
import axios from 'axios';

import Home from './pages/Home';

Object.assign(window, { axios });

ReactDOM.render(
    <React.Fragment>
        <CssBaseline />
        <Home />
    </React.Fragment>,
    document.getElementById('react-root')
);
