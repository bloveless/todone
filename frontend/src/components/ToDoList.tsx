import * as React from 'react';
import { createStyles, withStyles, WithStyles, Theme } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListSubheader from '@material-ui/core/ListSubheader';
import ToDoListItem from './ToDoListItem';

const styles = (theme: Theme) => createStyles({
    root: {
        width: '100%',
        margin: 'auto',
        backgroundColor: theme.palette.background.paper,
    },
});

export interface ToDo {
    id: string;
    text: string;
    complete: boolean;
};
interface ToDoListProps extends WithStyles<typeof styles> { };
interface ToDoListState {
    todos: ToDo[];
    activeId: string;
};

class ToDoList extends React.Component<ToDoListProps, ToDoListState> {
    state = {
        todos: [{
            id: '41ed2588-3c7e-4ecc-809e-23ddd5b017e0',
            text: 'First ToDo',
            complete: true,
        }, {
            id: 'f0903fa1-bf65-4513-82f8-62157b42d016',
            text: 'Second ToDo',
            complete: false,
        }],
        activeId: '',
    };

    handleToggleComplete = (todoId: string) => () => {
        const { todos } = this.state;
        todos.forEach((currentTodo) => {
            if (currentTodo.id === todoId) {
                currentTodo.complete = !currentTodo.complete;
            }
        });

        this.setState({
            todos
        });
    };

    handleToggleActive = (todoId: string) => () => {
        const { activeId } = this.state;

        this.setState({
            activeId: (todoId !== activeId) ? todoId : '',
        });
    }

    render() {
        const { classes } = this.props;
        const { todos, activeId } = this.state;

        return (
            <List className={classes.root}>
                {todos.map((todo) => (
                    <ToDoListItem
                        key={todo.id}
                        todo={todo}
                        active={activeId === todo.id}
                        handleToggleComplete={this.handleToggleComplete}
                        handleToggleActive={this.handleToggleActive}
                    />
                ))}
            </List>
        );
    }
}

export default withStyles(styles)(ToDoList);
