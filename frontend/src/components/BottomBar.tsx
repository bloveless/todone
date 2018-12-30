import * as React from 'react';
import { WithStyles, createStyles, withStyles } from '@material-ui/core/styles';
import BottomNavigation from '@material-ui/core/BottomNavigation';
import BottomNavigationAction from '@material-ui/core/BottomNavigationAction';
import LensIcon from '@material-ui/icons/LensOutlined';
import CheckCircleIcon from '@material-ui/icons/CheckCircleOutlined';

const styles = createStyles({
    root: {
    },
});

interface BottomBarProps extends WithStyles<typeof styles> {};
interface BottomBarState {
    value: number;
}

class BottomBar extends React.Component<BottomBarProps, BottomBarState> {
    state = {
        value: 0,
    };

    handleChange = (event: React.ChangeEvent<HTMLInputElement>, value: number) => {
        this.setState({ value });
    };

    render() {
        const { classes } = this.props;
        const { value } = this.state;

        return (
            <BottomNavigation
                value={value}
                onChange={this.handleChange}
                showLabels
                className={classes.root}
            >
                <BottomNavigationAction label="Incomplete" icon={<LensIcon />} />
                <BottomNavigationAction label="Complete" icon={<CheckCircleIcon />} />
            </BottomNavigation>
        );
    }
}

export default withStyles(styles)(BottomBar);
