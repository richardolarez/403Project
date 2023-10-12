=begin comment

Swagger Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

OpenAPI spec version: 1.0.0
Contact: apiteam@swagger.io
Generated by: https://github.com/swagger-api/swagger-codegen.git

=end comment

=cut

#
# NOTE: This class is auto generated by the swagger code generator program. 
# Do not edit the class manually.
# Ref: https://github.com/swagger-api/swagger-codegen
#
package WWW::SwaggerClient::ApiFactory;

use strict;
use warnings;
use utf8;

use Carp;
use Module::Find;

usesub WWW::SwaggerClient::Object;

use WWW::SwaggerClient::ApiClient;

=head1 Name

	WWW::SwaggerClient::ApiFactory - constructs APIs to retrieve WWW::SwaggerClient objects

=head1 Synopsis

	package My::Petstore::App;
	
	use WWW::SwaggerClient::ApiFactory;
	
	my $api_factory = WWW::SwaggerClient::ApiFactory->new( ... ); # any args for ApiClient constructor
							  
	# later...
	my $pet_api = $api_factory->get_api('Pet');  
	
	# $pet_api isa WWW::SwaggerClient::PetApi
	
	my $pet = $pet_api->get_pet_by_id(pet_id => $pet_id);
	
	# object attributes have proper accessors:
	printf "Pet's name is %s", $pet->name;
	
	# change the value stored on the object:
	$pet->name('Dave'); 

=cut

# Load all the API classes and construct a lookup table at startup time
my %_apis = map { $_ =~ /^WWW::SwaggerClient::(.*)$/; $1 => $_ } 
			grep {$_ =~ /Api$/} 
			usesub 'WWW::SwaggerClient';

=head1 new($api_client)
	
	create a new WWW::SwaggerClient::ApiFactory instance with the given WWW::SwaggerClient::ApiClient instance.

=head1 new(%parameters)

	Any parameters are optional, and are passed to and stored on the api_client object.
	
	See L<WWW::SwaggerClient::ApiClient> and L<WWW::SwaggerClient::Configuration> for valid parameters

=cut	

sub new {
    my ($class) = shift;

    my $api_client;
    if ($_[0] && ref $_[0] && ref $_[0] eq 'WWW::SwaggerClient::ApiClient' ) {
        $api_client = $_[0];
    } else {
        $api_client = WWW::SwaggerClient::ApiClient->new(@_);
    }
    bless { api_client => $api_client }, $class;
}

=head1 get_api($which)

	Returns an API object of the requested type. 
	
	$which is a nickname for the class: 
	
		FooBarClient::BazApi has nickname 'Baz'
		
=cut

sub get_api {
	my ($self, $which) = @_;
	croak "API not specified" unless $which;
	my $api_class = $_apis{"${which}Api"} || croak "No known API for '$which'";
	return $api_class->new($self->api_client); 
}

=head1 api_client()

	Returns the api_client object, should you ever need it.
	
=cut

sub api_client { $_[0]->{api_client} }

=head1 apis_available()
=cut 

sub apis_available { return map { $_ =~ s/Api$//; $_ } sort keys %_apis }

=head1 classname_for()
=cut

sub classname_for {
	my ($self, $api_name) = @_;
	return $_apis{"${api_name}Api"};
}


1;
