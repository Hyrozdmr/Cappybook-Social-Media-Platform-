
const Post = (props) => {
  return (
      <article key={props.post._id}>
        <p>{props.post.message}</p>
        <p>Likes: {props.post.likes}</p>
      </article>
  );
};

export default Post;
