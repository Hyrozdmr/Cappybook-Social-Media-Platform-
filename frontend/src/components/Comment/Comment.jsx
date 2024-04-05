import userImage from "/src/static/img/user_image.png";
import "./Comment.css"
import icon from "/src/static/img/x-button.png"

const Comment= ({comment, onDelete}) => {

const handleDeleteCommentClick = () => {
    onDelete(comment._id);
};

return (
    <div className="comment-info delete-comment">
        <div className="comment-user">
            <div>
            <img className="comment-user-image" src={userImage} alt="image" />
            </div>
            <div className="comment-message">
            <h6>Capybara</h6>
            <p>{comment.message}</p>
            </div>
        </div>
        <div className="delete-comment-button">
            <img className="delete-button" src={icon} onClick={handleDeleteCommentClick} />
        </div>

    </div>
    
);
};

export default Comment;