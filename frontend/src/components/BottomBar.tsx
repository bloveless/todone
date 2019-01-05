import * as React from 'react';
import { WithStyles, createStyles, withStyles } from '@material-ui/core/styles';
import BottomNavigation from '@material-ui/core/BottomNavigation';
import BottomNavigationAction from '@material-ui/core/BottomNavigationAction';
import CheckBox from '@material-ui/icons/CheckBox';
import CheckBoxOutlineBlank from '@material-ui/icons/CheckBoxOutlineBlank';

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
                <BottomNavigationAction label="Incomplete" icon={<CheckBoxOutlineBlank />} />
                <BottomNavigationAction label="Complete" icon={<CheckBox />} />
            </BottomNavigation>
        );
    }
}

export default withStyles(styles)(BottomBar);
