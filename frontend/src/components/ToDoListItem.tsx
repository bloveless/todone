import * as React from 'react';
import { ToDo } from './ToDoList';
import Checkbox from '@material-ui/core/Checkbox';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import Switch from '@material-ui/core/Switch';
import Tooltip from '@material-ui/core/Tooltip';
import IconButton from '@material-ui/core/IconButton';
import MoreVertIcon from '@material-ui/icons/MoreVert';

interface ToDoListItemProps {
    todo: ToDo;
    active: boolean;
    handleToggleComplete(todoId: string): () => void;
    handleToggleActive(todoId: string): () => void;
};

interface ToDoListItemState {
    subMenuVisible: boolean;
}

class ToDoListItem extends React.Component<ToDoListItemProps, ToDoListItemState> {
    state = {
        subMenuVisible: false,
    };

    render() {
        const { todo, active, handleToggleComplete, handleToggleActive } = this.props;
        const { subMenuVisible } = this.state;

        return (
            <React.Fragment>
                <ListItem divider>
                    <Tooltip title="Toggle complete" enterDelay={1000}>
                        <Checkbox
                            checked={todo.complete}
                            tabIndex={-1}
                            onClick={handleToggleComplete(todo.id)}
                        />
                    </Tooltip>
                    <ListItemText primary={todo.text} />
                    <Tooltip title="Toggle currently active" enterDelay={1000}>
                        <Switch
                            onChange={handleToggleActive(todo.id)}
                            checked={active && !todo.complete}
                            disabled={todo.complete}
                        />
                    </Tooltip>
                    <IconButton
                        aria-label="More"
                        aria-owns={open ? 'long-menu' : undefined}
                        aria-haspopup="true"
                        onClick={() => (this.setState({ subMenuVisible: !this.state.subMenuVisible }))}
                    >
                        <MoreVertIcon />
                    </IconButton>
                </ListItem>
                {subMenuVisible && (
                    <ListItem divider>
                        <ListItemText inset primary={"10:00 AM - 10:15 AM"} />
                    </ListItem>
                )}
            </React.Fragment>
        );
    }
}

export default ToDoListItem;
