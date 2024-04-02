const Comment = ({ comment, user }) => {
  return (
    <article key={comment._id}>
      <div>{comment.message}</div>
      <div>User ID: {user._id}</div>
    </article>
  );
};

export default Comment;

