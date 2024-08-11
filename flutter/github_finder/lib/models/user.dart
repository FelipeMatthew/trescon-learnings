class User {
  final String name;
  final String avatar_url;
  final int followers;
  final int public_repos;

  User(this.name, this.avatar_url, this.followers, this.public_repos);

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      json['name'],
      json['avatar_url'],
      json['followers'],
      json['public_repos'],
    );
  }
}
