import * as React from 'react';
import TopBar from '../components/TopBar';
import ToDoList from '../components/ToDoList';
import BottomBar from '../components/BottomBar';
import Grid from '@material-ui/core/Grid';

class Home extends React.Component<{}, {}> {
    render() {
        return (
            <React.Fragment>
                <TopBar />
                <div style={{ backgroundColor: 'white' }}>
                    <Grid
                        container
                        direction="row"
                        justify="center"
                        alignItems="center"
                    >
                        <Grid item xs={12} md={6}>
                            <ToDoList />
                        </Grid>
                    </Grid>
                </div>
                <BottomBar />
            </React.Fragment>
        );
    }
}

export default Home;
