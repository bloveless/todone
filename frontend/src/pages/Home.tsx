import * as React from 'react';
import {createStyles, withStyles, WithStyles, Theme} from '@material-ui/core/styles';
import TopBar from '../components/TopBar';
import ToDoList from '../components/ToDoList';
import BottomBar from '../components/BottomBar';

const styles = (theme: Theme) => createStyles({
    root: {
        width: '100%',
        margin: 'auto',
        backgroundColor: theme.palette.background.paper,
    },
});

interface HomeProps extends WithStyles<typeof styles> {
}

class Home extends React.Component<HomeProps, {}> {
    render() {
        const {classes} = this.props;

        return (
            <div className={classes.root}>
                <TopBar/>
                <ToDoList/>
                <BottomBar/>
            </div>
        );
    }
}

export default withStyles(styles)(Home);
