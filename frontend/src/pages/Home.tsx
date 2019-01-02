import * as React from 'react';
import { createStyles, withStyles, WithStyles, Theme } from '@material-ui/core/styles';
import TopBar from '../components/TopBar';
import ToDoList from '../components/ToDoList';
import BottomBar from '../components/BottomBar';
import Grid from '@material-ui/core/Grid';

const styles = (theme: Theme) => createStyles({
    root: {
        width: '100%',
        margin: 'auto',
        backgroundColor: theme.palette.background.paper,
    },
});

interface HomeProps extends WithStyles<typeof styles> {};

class Home extends React.Component<HomeProps, {}> {
    render() {
        const { classes } = this.props;

        return (
            <div className={classes.root}>
                <TopBar />
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
                <BottomBar />
            </div>
        );
    }
}

export default withStyles(styles)(Home);
